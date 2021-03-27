// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StringResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsStringResponse(buf []byte, offset flatbuffers.UOffsetT) *StringResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StringResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *StringResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StringResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StringResponse) Output() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *StringResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StringResponseAddOutput(builder *flatbuffers.Builder, output flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(output), 0)
}
func StringResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func StringResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
