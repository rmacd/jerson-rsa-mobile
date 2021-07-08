// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import "strconv"

type Hash int32

const (
	HashMD5    Hash = 0
	HashSHA1   Hash = 1
	HashSHA224 Hash = 2
	HashSHA256 Hash = 3
	HashSHA384 Hash = 4
	HashSHA512 Hash = 5
)

var EnumNamesHash = map[Hash]string{
	HashMD5:    "MD5",
	HashSHA1:   "SHA1",
	HashSHA224: "SHA224",
	HashSHA256: "SHA256",
	HashSHA384: "SHA384",
	HashSHA512: "SHA512",
}

var EnumValuesHash = map[string]Hash{
	"MD5":    HashMD5,
	"SHA1":   HashSHA1,
	"SHA224": HashSHA224,
	"SHA256": HashSHA256,
	"SHA384": HashSHA384,
	"SHA512": HashSHA512,
}

func (v Hash) String() string {
	if s, ok := EnumNamesHash[v]; ok {
		return s
	}
	return "Hash(" + strconv.FormatInt(int64(v), 10) + ")"
}
