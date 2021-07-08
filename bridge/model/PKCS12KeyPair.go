// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PKCS12KeyPair struct {
	_tab flatbuffers.Table
}

func GetRootAsPKCS12KeyPair(buf []byte, offset flatbuffers.UOffsetT) *PKCS12KeyPair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PKCS12KeyPair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPKCS12KeyPair(buf []byte, offset flatbuffers.UOffsetT) *PKCS12KeyPair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PKCS12KeyPair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PKCS12KeyPair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PKCS12KeyPair) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PKCS12KeyPair) PrivateKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PKCS12KeyPair) PublicKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PKCS12KeyPair) Certificate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PKCS12KeyPairStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PKCS12KeyPairAddPrivateKey(builder *flatbuffers.Builder, privateKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(privateKey), 0)
}
func PKCS12KeyPairAddPublicKey(builder *flatbuffers.Builder, publicKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(publicKey), 0)
}
func PKCS12KeyPairAddCertificate(builder *flatbuffers.Builder, certificate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(certificate), 0)
}
func PKCS12KeyPairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
