// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import "strconv"

type SaltLength int32

const (
	SaltLengthAUTO        SaltLength = 0
	SaltLengthEQUALS_HASH SaltLength = 1
)

var EnumNamesSaltLength = map[SaltLength]string{
	SaltLengthAUTO:        "AUTO",
	SaltLengthEQUALS_HASH: "EQUALS_HASH",
}

var EnumValuesSaltLength = map[string]SaltLength{
	"AUTO":        SaltLengthAUTO,
	"EQUALS_HASH": SaltLengthEQUALS_HASH,
}

func (v SaltLength) String() string {
	if s, ok := EnumNamesSaltLength[v]; ok {
		return s
	}
	return "SaltLength(" + strconv.FormatInt(int64(v), 10) + ")"
}
