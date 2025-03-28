// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/19 18:51:50                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base                                                                                                        *
// * File: base.go                                                                                                     *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base

import (
	"encoding/binary"
	"unsafe"
)

const (
	MinUseRune byte = ' '
	MaxUseRune byte = '~'
	StdPadding rune = '='
	NotPadding rune = -1
)

type IEncoding interface {
	Encode(src []byte) ([]byte, error)
	Decode(src []byte) ([]byte, error)
}

func StringToBytes(src string) []byte {
	return unsafe.Slice(unsafe.StringData(src), len(src))
}

func BytesToString(src []byte) string {
	return unsafe.String(&src[0], len(src))
}

func HasRepeatElem[T comparable](array []T) bool {
	mp := make(map[T]struct{})
	for _, v := range array {
		mp[v] = struct{}{}
	}
	return len(mp) != len(array)
}

func HasRepeatChar(characters string) bool {
	mp := make(map[rune]struct{})
	for _, v := range characters {
		mp[v] = struct{}{}
	}
	return len(mp) != len(characters)
}

func IsIllegalCharacter(c rune) bool {
	if c == NotPadding {
		return false
	}
	return c < rune(MinUseRune) || c > rune(MaxUseRune)
}

func TrimNewLines(src []byte) []byte {
	dst := make([]byte, len(src))
	offset := 0
	for _, v := range src {
		if v == '\r' || v == '\n' {
			continue
		}
		dst[offset] = v
		offset++
	}
	return dst[:offset]
}

func IsNewLineChar(c byte) bool {
	return c == '\r' || c == '\n'
}

func Uint16ToBytes(in uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, in)
	return bytes
}
