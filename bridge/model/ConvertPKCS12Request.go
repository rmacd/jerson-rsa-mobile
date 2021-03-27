// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConvertPKCS12Request struct {
	_tab flatbuffers.Table
}

func GetRootAsConvertPKCS12Request(buf []byte, offset flatbuffers.UOffsetT) *ConvertPKCS12Request {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConvertPKCS12Request{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ConvertPKCS12Request) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConvertPKCS12Request) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConvertPKCS12Request) Pkcs12() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ConvertPKCS12Request) Password() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ConvertPKCS12RequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ConvertPKCS12RequestAddPkcs12(builder *flatbuffers.Builder, pkcs12 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pkcs12), 0)
}
func ConvertPKCS12RequestAddPassword(builder *flatbuffers.Builder, password flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(password), 0)
}
func ConvertPKCS12RequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
