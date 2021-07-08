// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Base64Request struct {
	_tab flatbuffers.Table
}

func GetRootAsBase64Request(buf []byte, offset flatbuffers.UOffsetT) *Base64Request {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Base64Request{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBase64Request(buf []byte, offset flatbuffers.UOffsetT) *Base64Request {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Base64Request{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Base64Request) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Base64Request) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Base64Request) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func Base64RequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Base64RequestAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(message), 0)
}
func Base64RequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
