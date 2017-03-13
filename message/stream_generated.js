// automatically generated by the FlatBuffers compiler, do not modify

/**
 * @const
 * @namespace
 */
var message = message || {};

/**
 * @enum
 */
message.InputEncoding = {
  Opus: 1,
  PCM: 2,
  AAC: 3
};

/**
 * @enum
 */
message.ResponseStatus = {
  OK: 1,
  NoStream: 2,
  NotAllowed: 3,
  Error: 4
};

/**
 * @enum
 */
message.StreamStatus = {
  CREATED: 0,
  STREAMING: 1,
  STOPPED: 2,
  ERROR: 3
};

/**
 * @enum
 */
message.Message = {
  NONE: 0,
  StreamBroadCast: 1,
  StreamPause: 2,
  StreamStop: 3,
  StreamFrame: 4,
  StreamComment: 5,
  StreamSubscribe: 6,
  StreamResponse: 7,
  StreamStatus: 8,
  StreamUnSubscribe: 9
};

/**
 * @constructor
 */
message.StreamSubscribe = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamSubscribe}
 */
message.StreamSubscribe.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamSubscribe=} obj
 * @returns {message.StreamSubscribe}
 */
message.StreamSubscribe.getRootAsStreamSubscribe = function(bb, obj) {
  return (obj || new message.StreamSubscribe).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns {message.StreamStatus}
 */
message.StreamSubscribe.prototype.status = function() {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? /** @type {message.StreamStatus} */ (this.bb.readInt8(this.bb_pos + offset)) : message.StreamStatus.CREATED;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamSubscribe.startStreamSubscribe = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {message.StreamStatus} status
 */
message.StreamSubscribe.addStatus = function(builder, status) {
  builder.addFieldInt8(0, status, message.StreamStatus.CREATED);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamSubscribe.endStreamSubscribe = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamUnSubscribe = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamUnSubscribe}
 */
message.StreamUnSubscribe.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamUnSubscribe=} obj
 * @returns {message.StreamUnSubscribe}
 */
message.StreamUnSubscribe.getRootAsStreamUnSubscribe = function(bb, obj) {
  return (obj || new message.StreamUnSubscribe).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamUnSubscribe.startStreamUnSubscribe = function(builder) {
  builder.startObject(0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamUnSubscribe.endStreamUnSubscribe = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamStatus = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamStatus}
 */
message.StreamStatus.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamStatus=} obj
 * @returns {message.StreamStatus}
 */
message.StreamStatus.getRootAsStreamStatus = function(bb, obj) {
  return (obj || new message.StreamStatus).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns {message.StreamStatus}
 */
message.StreamStatus.prototype.status = function() {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? /** @type {message.StreamStatus} */ (this.bb.readInt8(this.bb_pos + offset)) : message.StreamStatus.CREATED;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamStatus.startStreamStatus = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {message.StreamStatus} status
 */
message.StreamStatus.addStatus = function(builder, status) {
  builder.addFieldInt8(0, status, message.StreamStatus.CREATED);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamStatus.endStreamStatus = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamBroadCast = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamBroadCast}
 */
message.StreamBroadCast.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamBroadCast=} obj
 * @returns {message.StreamBroadCast}
 */
message.StreamBroadCast.getRootAsStreamBroadCast = function(bb, obj) {
  return (obj || new message.StreamBroadCast).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns {message.InputEncoding}
 */
message.StreamBroadCast.prototype.encoding = function() {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? /** @type {message.InputEncoding} */ (this.bb.readInt8(this.bb_pos + offset)) : message.InputEncoding.Opus;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamBroadCast.startStreamBroadCast = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {message.InputEncoding} encoding
 */
message.StreamBroadCast.addEncoding = function(builder, encoding) {
  builder.addFieldInt8(0, encoding, message.InputEncoding.Opus);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamBroadCast.endStreamBroadCast = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamPause = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamPause}
 */
message.StreamPause.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamPause=} obj
 * @returns {message.StreamPause}
 */
message.StreamPause.getRootAsStreamPause = function(bb, obj) {
  return (obj || new message.StreamPause).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamPause.startStreamPause = function(builder) {
  builder.startObject(0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamPause.endStreamPause = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamStop = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamStop}
 */
message.StreamStop.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamStop=} obj
 * @returns {message.StreamStop}
 */
message.StreamStop.getRootAsStreamStop = function(bb, obj) {
  return (obj || new message.StreamStop).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamStop.startStreamStop = function(builder) {
  builder.startObject(0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamStop.endStreamStop = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamFrame = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamFrame}
 */
message.StreamFrame.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamFrame=} obj
 * @returns {message.StreamFrame}
 */
message.StreamFrame.getRootAsStreamFrame = function(bb, obj) {
  return (obj || new message.StreamFrame).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns {number}
 */
message.StreamFrame.prototype.frameSize = function() {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.readUint8(this.bb_pos + offset) : 0;
};

/**
 * @returns {number}
 */
message.StreamFrame.prototype.sampleRate = function() {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? this.bb.readUint32(this.bb_pos + offset) : 0;
};

/**
 * @returns {number}
 */
message.StreamFrame.prototype.channels = function() {
  var offset = this.bb.__offset(this.bb_pos, 8);
  return offset ? this.bb.readUint8(this.bb_pos + offset) : 0;
};

/**
 * @param {number} index
 * @returns {number}
 */
message.StreamFrame.prototype.frame = function(index) {
  var offset = this.bb.__offset(this.bb_pos, 10);
  return offset ? this.bb.readUint8(this.bb.__vector(this.bb_pos + offset) + index) : 0;
};

/**
 * @returns {number}
 */
message.StreamFrame.prototype.frameLength = function() {
  var offset = this.bb.__offset(this.bb_pos, 10);
  return offset ? this.bb.__vector_len(this.bb_pos + offset) : 0;
};

/**
 * @returns {Uint8Array}
 */
message.StreamFrame.prototype.frameArray = function() {
  var offset = this.bb.__offset(this.bb_pos, 10);
  return offset ? new Uint8Array(this.bb.bytes().buffer, this.bb.bytes().byteOffset + this.bb.__vector(this.bb_pos + offset), this.bb.__vector_len(this.bb_pos + offset)) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamFrame.startStreamFrame = function(builder) {
  builder.startObject(4);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {number} frameSize
 */
message.StreamFrame.addFrameSize = function(builder, frameSize) {
  builder.addFieldInt8(0, frameSize, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {number} sampleRate
 */
message.StreamFrame.addSampleRate = function(builder, sampleRate) {
  builder.addFieldInt32(1, sampleRate, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {number} channels
 */
message.StreamFrame.addChannels = function(builder, channels) {
  builder.addFieldInt8(2, channels, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} frameOffset
 */
message.StreamFrame.addFrame = function(builder, frameOffset) {
  builder.addFieldOffset(3, frameOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {Array.<number>} data
 * @returns {flatbuffers.Offset}
 */
message.StreamFrame.createFrameVector = function(builder, data) {
  builder.startVector(1, data.length, 1);
  for (var i = data.length - 1; i >= 0; i--) {
    builder.addInt8(data[i]);
  }
  return builder.endVector();
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {number} numElems
 */
message.StreamFrame.startFrameVector = function(builder, numElems) {
  builder.startVector(1, numElems, 1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamFrame.endStreamFrame = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamComment = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamComment}
 */
message.StreamComment.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamComment=} obj
 * @returns {message.StreamComment}
 */
message.StreamComment.getRootAsStreamComment = function(bb, obj) {
  return (obj || new message.StreamComment).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
message.StreamComment.prototype.userName = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
message.StreamComment.prototype.text = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamComment.startStreamComment = function(builder) {
  builder.startObject(2);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} userNameOffset
 */
message.StreamComment.addUserName = function(builder, userNameOffset) {
  builder.addFieldOffset(0, userNameOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} textOffset
 */
message.StreamComment.addText = function(builder, textOffset) {
  builder.addFieldOffset(1, textOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamComment.endStreamComment = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamResponse = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamResponse}
 */
message.StreamResponse.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamResponse=} obj
 * @returns {message.StreamResponse}
 */
message.StreamResponse.getRootAsStreamResponse = function(bb, obj) {
  return (obj || new message.StreamResponse).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @returns {message.ResponseStatus}
 */
message.StreamResponse.prototype.status = function() {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? /** @type {message.ResponseStatus} */ (this.bb.readInt8(this.bb_pos + offset)) : message.ResponseStatus.OK;
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamResponse.startStreamResponse = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {message.ResponseStatus} status
 */
message.StreamResponse.addStatus = function(builder, status) {
  builder.addFieldInt8(0, status, message.ResponseStatus.OK);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamResponse.endStreamResponse = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
message.StreamMessage = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {message.StreamMessage}
 */
message.StreamMessage.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {message.StreamMessage=} obj
 * @returns {message.StreamMessage}
 */
message.StreamMessage.getRootAsStreamMessage = function(bb, obj) {
  return (obj || new message.StreamMessage).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
message.StreamMessage.prototype.eventId = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
message.StreamMessage.prototype.streamId = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @returns {message.Message}
 */
message.StreamMessage.prototype.messageType = function() {
  var offset = this.bb.__offset(this.bb_pos, 8);
  return offset ? /** @type {message.Message} */ (this.bb.readUint8(this.bb_pos + offset)) : message.Message.NONE;
};

/**
 * @param {flatbuffers.Table} obj
 * @returns {?flatbuffers.Table}
 */
message.StreamMessage.prototype.message = function(obj) {
  var offset = this.bb.__offset(this.bb_pos, 10);
  return offset ? this.bb.__union(obj, this.bb_pos + offset) : null;
};

/**
 * @returns {flatbuffers.Long}
 */
message.StreamMessage.prototype.timestamp = function() {
  var offset = this.bb.__offset(this.bb_pos, 12);
  return offset ? this.bb.readInt64(this.bb_pos + offset) : this.bb.createLong(0, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 */
message.StreamMessage.startStreamMessage = function(builder) {
  builder.startObject(5);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} eventIdOffset
 */
message.StreamMessage.addEventId = function(builder, eventIdOffset) {
  builder.addFieldOffset(0, eventIdOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} streamIdOffset
 */
message.StreamMessage.addStreamId = function(builder, streamIdOffset) {
  builder.addFieldOffset(1, streamIdOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {message.Message} messageType
 */
message.StreamMessage.addMessageType = function(builder, messageType) {
  builder.addFieldInt8(2, messageType, message.Message.NONE);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} messageOffset
 */
message.StreamMessage.addMessage = function(builder, messageOffset) {
  builder.addFieldOffset(3, messageOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Long} timestamp
 */
message.StreamMessage.addTimestamp = function(builder, timestamp) {
  builder.addFieldInt64(4, timestamp, builder.createLong(0, 0));
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
message.StreamMessage.endStreamMessage = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} offset
 */
message.StreamMessage.finishStreamMessageBuffer = function(builder, offset) {
  builder.finish(offset);
};

// Exports for Node.js and RequireJS
this.message = message;
