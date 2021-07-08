// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DecryptOAEPBytesRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsDecryptOAEPBytesRequest(buf []byte, offset flatbuffers.UOffsetT) *DecryptOAEPBytesRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DecryptOAEPBytesRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDecryptOAEPBytesRequest(buf []byte, offset flatbuffers.UOffsetT) *DecryptOAEPBytesRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DecryptOAEPBytesRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DecryptOAEPBytesRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DecryptOAEPBytesRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DecryptOAEPBytesRequest) Ciphertext(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DecryptOAEPBytesRequest) CiphertextLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DecryptOAEPBytesRequest) CiphertextBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DecryptOAEPBytesRequest) MutateCiphertext(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *DecryptOAEPBytesRequest) Label() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DecryptOAEPBytesRequest) Hash() Hash {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Hash(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DecryptOAEPBytesRequest) MutateHash(n Hash) bool {
	return rcv._tab.MutateInt32Slot(8, int32(n))
}

func (rcv *DecryptOAEPBytesRequest) PrivateKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DecryptOAEPBytesRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DecryptOAEPBytesRequestAddCiphertext(builder *flatbuffers.Builder, ciphertext flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ciphertext), 0)
}
func DecryptOAEPBytesRequestStartCiphertextVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DecryptOAEPBytesRequestAddLabel(builder *flatbuffers.Builder, label flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(label), 0)
}
func DecryptOAEPBytesRequestAddHash(builder *flatbuffers.Builder, hash Hash) {
	builder.PrependInt32Slot(2, int32(hash), 0)
}
func DecryptOAEPBytesRequestAddPrivateKey(builder *flatbuffers.Builder, privateKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(privateKey), 0)
}
func DecryptOAEPBytesRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
