// Copyright © 2021 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package verifier

import (
	"crypto/elliptic"
	"errors"
	"math/big"

	"github.com/getamis/alice/crypto/birkhoffinterpolation"
	"github.com/getamis/alice/crypto/ecpointgrouplaw"
	"github.com/getamis/alice/crypto/oprf"
	"github.com/getamis/alice/crypto/tss"
	"github.com/getamis/alice/crypto/zkproof"
	"github.com/getamis/alice/internal/message/types"
	"github.com/getamis/sirius/log"
)

var (
	ErrIndentityPublicKey    = errors.New("identity public key")
	ErrNoSelfBk              = errors.New("no self bk")
	ErrInconsistentPublicKey = errors.New("inconsistent public key")
)

type userHandler0 struct {
	peerManager types.PeerManager
	publicKey   *ecpointgrouplaw.ECPoint
	peers       map[string]*peer
	bks         map[string]*birkhoffinterpolation.BkParameter
	curve       elliptic.Curve

	passwordRequester *oprf.Requester

	share           *big.Int
	serverGVerifier *zkproof.InteractiveSchnorrVerifier
	shareGProver    *zkproof.InteractiveSchnorrProver
}

func newUserHandler0(publicKey *ecpointgrouplaw.ECPoint, peerManager types.PeerManager, bks map[string]*birkhoffinterpolation.BkParameter, password []byte) (*userHandler0, error) {
	if publicKey.IsIdentity() {
		return nil, ErrIndentityPublicKey
	}

	// Construct peers
	curve := publicKey.GetCurve()
	peers, err := buildPeers(peerManager.SelfID(), curve.Params().N, bks)
	if err != nil {
		return nil, err
	}

	// Build requesters
	requester, err := oprf.NewRequester(password)
	if err != nil {
		return nil, err
	}

	return &userHandler0{
		publicKey:   publicKey,
		peerManager: peerManager,
		peers:       peers,
		bks:         bks,
		curve:       curve,

		passwordRequester: requester,
	}, nil
}

func (p *userHandler0) MessageType() types.MessageType {
	return types.MessageType(Type_MsgServer0)
}

func (p *userHandler0) GetRequiredMessageCount() uint32 {
	return 1
}

func (p *userHandler0) IsHandled(logger log.Logger, id string) bool {
	peer, ok := p.peers[id]
	if !ok {
		logger.Debug("Peer not found")
		return false
	}
	return peer.GetMessage(p.MessageType()) != nil
}

func (p *userHandler0) HandleMessage(logger log.Logger, message types.Message) error {
	msg := getMessage(message)
	server0 := msg.GetServer0()
	id := msg.GetId()
	peer, ok := p.peers[id]
	if !ok {
		logger.Debug("Peer not found")
		return tss.ErrPeerNotFound
	}

	// Compute shares
	var err error
	p.share, err = p.passwordRequester.Compute(server0.PasswordResponse)
	if err != nil {
		logger.Debug("Failed to compute old share", "err", err)
		return err
	}
	p.shareGProver, err = zkproof.NewInteractiveSchnorrProver(p.share, p.curve)
	if err != nil {
		logger.Debug("Failed to create old share prover", "err", err)
		return err
	}
	// Compute server g and build its verifier
	p.serverGVerifier, err = zkproof.NewInteractiveSchnorrVerifier(server0.ServerGProver1)
	if err != nil {
		logger.Debug("Failed to new server g verifier", "err", err)
		return err
	}
	// Send to Server
	p.peerManager.MustSend(message.GetId(), &Message{
		Type: Type_MsgUser1,
		Id:   p.peerManager.SelfID(),
		Body: &Message_User1{
			User1: &BodyUser1{
				ShareGProver1:    p.shareGProver.GetInteractiveSchnorrProver1Message(),
				ServerGVerifier1: p.serverGVerifier.GetInteractiveSchnorrVerifier1Message()},
		},
	})
	return peer.AddMessage(msg)
}

func (p *userHandler0) Finalize(logger log.Logger) (types.Handler, error) {
	return newUserHandler1(p)
}

func (p *userHandler0) GetFirstMessage() *Message {
	return &Message{
		Type: Type_MsgUser0,
		Id:   p.peerManager.SelfID(),
		Body: &Message_User0{
			User0: &BodyUser0{
				PasswordRequest: p.passwordRequester.GetRequestMessage(),
			},
		},
	}
}

func buildPeers(selfId string, fieldOrder *big.Int, bks map[string]*birkhoffinterpolation.BkParameter) (map[string]*peer, error) {
	lenBks := len(bks)

	// Ensure self bk exists
	_, ok := bks[selfId]
	if !ok {
		return nil, ErrNoSelfBk
	}

	// Compute coefficients
	allBKs := make(birkhoffinterpolation.BkParameters, lenBks)
	peers := make([]*peer, lenBks)
	i := 0
	for id, bk := range bks {
		allBKs[i] = bk
		peers[i] = newPeer(id, bk)
		i++
	}
	gots, err := allBKs.ComputeBkCoefficient(tss.PasswordThreshold, fieldOrder)
	if err != nil {
		return nil, err
	}
	for i, co := range gots {
		peers[i].bkCoefficient = co
	}

	// Build peer map
	peerMaps := make(map[string]*peer, lenBks)
	for _, p := range peers {
		peerMaps[p.Id] = p
	}
	return peerMaps, nil
}

func validatePubKey(serverCo *big.Int, serverShareG *ecpointgrouplaw.ECPoint, userCo *big.Int, userShareG *ecpointgrouplaw.ECPoint, pubkey *ecpointgrouplaw.ECPoint) error {
	return birkhoffinterpolation.ValidatePublicKeyWithBkCoefficients([]*big.Int{
		serverCo,
		userCo,
	}, []*ecpointgrouplaw.ECPoint{
		serverShareG,
		userShareG,
	}, pubkey)
}