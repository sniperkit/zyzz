// automatically generated by the FlatBuffers compiler, do not modify

package message

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StreamResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsStreamResponse(buf []byte, offset flatbuffers.UOffsetT) *StreamResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StreamResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *StreamResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StreamResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StreamResponse) Status() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *StreamResponse) MutateStatus(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func StreamResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StreamResponseAddStatus(builder *flatbuffers.Builder, status int8) {
	builder.PrependInt8Slot(0, status, 1)
}
func StreamResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}