/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.birkhoffinterpolation.BkParameterMessage', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.birkhoffinterpolation.BkParameterMessage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.birkhoffinterpolation.BkParameterMessage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.birkhoffinterpolation.BkParameterMessage.displayName = 'proto.birkhoffinterpolation.BkParameterMessage';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.toObject = function(opt_includeInstance) {
  return proto.birkhoffinterpolation.BkParameterMessage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.birkhoffinterpolation.BkParameterMessage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.birkhoffinterpolation.BkParameterMessage.toObject = function(includeInstance, msg) {
  var f, obj = {
    x: msg.getX_asB64(),
    rank: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.birkhoffinterpolation.BkParameterMessage}
 */
proto.birkhoffinterpolation.BkParameterMessage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.birkhoffinterpolation.BkParameterMessage;
  return proto.birkhoffinterpolation.BkParameterMessage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.birkhoffinterpolation.BkParameterMessage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.birkhoffinterpolation.BkParameterMessage}
 */
proto.birkhoffinterpolation.BkParameterMessage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setX(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setRank(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.birkhoffinterpolation.BkParameterMessage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.birkhoffinterpolation.BkParameterMessage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.birkhoffinterpolation.BkParameterMessage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getX_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getRank();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
};


/**
 * optional bytes x = 1;
 * @return {!(string|Uint8Array)}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.getX = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes x = 1;
 * This is a type-conversion wrapper around `getX()`
 * @return {string}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.getX_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getX()));
};


/**
 * optional bytes x = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getX()`
 * @return {!Uint8Array}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.getX_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getX()));
};


/** @param {!(string|Uint8Array)} value */
proto.birkhoffinterpolation.BkParameterMessage.prototype.setX = function(value) {
  jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional uint32 rank = 2;
 * @return {number}
 */
proto.birkhoffinterpolation.BkParameterMessage.prototype.getRank = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.birkhoffinterpolation.BkParameterMessage.prototype.setRank = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


goog.object.extend(exports, proto.birkhoffinterpolation);