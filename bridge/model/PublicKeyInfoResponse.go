// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PublicKeyInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsPublicKeyInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *PublicKeyInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PublicKeyInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PublicKeyInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PublicKeyInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PublicKeyInfoResponse) Output(obj *PublicKeyInfo) *PublicKeyInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PublicKeyInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PublicKeyInfoResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PublicKeyInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PublicKeyInfoResponseAddOutput(builder *flatbuffers.Builder, output flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(output), 0)
}
func PublicKeyInfoResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func PublicKeyInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
