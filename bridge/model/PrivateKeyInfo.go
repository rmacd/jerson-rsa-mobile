// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PrivateKeyInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsPrivateKeyInfo(buf []byte, offset flatbuffers.UOffsetT) *PrivateKeyInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PrivateKeyInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPrivateKeyInfo(buf []byte, offset flatbuffers.UOffsetT) *PrivateKeyInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PrivateKeyInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PrivateKeyInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PrivateKeyInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PrivateKeyInfo) BitLen() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrivateKeyInfo) MutateBitLen(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *PrivateKeyInfo) Size() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrivateKeyInfo) MutateSize(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *PrivateKeyInfo) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PrivateKeyInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PrivateKeyInfoAddBitLen(builder *flatbuffers.Builder, bitLen int64) {
	builder.PrependInt64Slot(0, bitLen, 0)
}
func PrivateKeyInfoAddSize(builder *flatbuffers.Builder, size int64) {
	builder.PrependInt64Slot(1, size, 0)
}
func PrivateKeyInfoAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(error), 0)
}
func PrivateKeyInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
