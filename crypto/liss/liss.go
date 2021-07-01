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

package liss

import (
	"errors"
	"math/big"

	bqForm "github.com/getamis/alice/crypto/binaryquadraticform"
	"github.com/getamis/alice/crypto/homo/cl"
	"github.com/getamis/alice/internal/message"
	"github.com/getamis/alice/internal/message/types"
)

var (
	ErrNotReady = errors.New("not ready")

	msgTypes = []types.MessageType{
		types.MessageType(Type_BqCommitment),
		types.MessageType(Type_BqDecommitment),
	}
)

type Result struct {
	PublicKey *cl.PublicKey
	// Per group, per user
	Users [][]map[string]*UserResult
}

type UserResult struct {
	Bq    *bqForm.BQuadraticForm
	Share *big.Int
}

type Liss struct {
	*message.MsgMain

	ih *bqCommitmentHandler
}

func NewLiss(peerManager types.PeerManager, configs GroupConfigs, listener types.StateChangedListener) (*Liss, error) {
	numPeers := peerManager.NumPeers()
	ih, err := newBqCommitmentHandler(peerManager, configs)
	if err != nil {
		return nil, err
	}
	return &Liss{
		ih: ih,
		MsgMain: message.NewMsgMain(peerManager.SelfID(),
			numPeers,
			listener,
			ih,
			msgTypes...,
		),
	}, nil
}

func (m *Liss) Start() {
	m.MsgMain.Start()
	m.ih.broadcast(m.ih.bqMsg)
}

func (m *Liss) GetResult() (*Result, error) {
	if m.GetState() != types.StateDone {
		return nil, ErrNotReady
	}

	h := m.GetHandler()
	rh, ok := h.(*bqDecommitmentHandler)
	if !ok {
		return nil, ErrNotReady
	}

	users := make([][]map[string]*UserResult, len(rh.configs))
	for i, config := range rh.configs {
		users[i] = make([]map[string]*UserResult, config.Users)
		for j := 0; j < config.Users; j++ {
			users[i][j] = make(map[string]*UserResult)
			for k := range rh.shares[i][j] {
				users[i][j][k] = &UserResult{
					Share: rh.shares[i][j][k],
					Bq:    rh.shareCommitments[i][j][k],
				}
			}
		}
	}

	return &Result{
		Users:     users,
		PublicKey: rh.publicKey,
	}, nil
}