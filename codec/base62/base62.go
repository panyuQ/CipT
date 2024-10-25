// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/17 21:20:22                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base62                                                                                                      *
// * File: base62.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base62

import (
	"CipT/codec/base"
	"math"
)

const (
	codec          = "base62"
	stdEncoder     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	stdEncoderSize = 62
)

var StdCodec, _ = NewCodec(stdEncoder)

type base62Encoder struct {
	encodeMap [62]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}
	b := &base62Encoder{decodeMap: make(map[byte]int, stdEncoderSize)}
	mp := make(map[rune]struct{})
	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
		}
		b.encodeMap[k] = byte(v)
		b.decodeMap[byte(v)] = k
		mp[v] = struct{}{}
	}
	if len(mp) != stdEncoderSize {
		return nil, base.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base62Encoder) maxEncodedLen(n int) int {
	return int(math.Ceil(math.Log(256) / math.Log(62) * float64(n)))
}

func (b *base62Encoder) encode(src []byte) []byte {
	rs := 0
	cs := b.maxEncodedLen(len(src))
	dst := make([]byte, cs)
	for k := range src {
		c := 0
		v := int(src[k])
		for j := cs - 1; j >= 0 && (v != 0 || c < rs); j-- {
			v += 256 * int(dst[j])
			dst[j] = byte(v % stdEncoderSize)
			v /= stdEncoderSize
			c++
		}
		rs = c
	}
	for k := range dst {
		dst[k] = b.encodeMap[dst[k]]
	}
	if cs > rs {
		return dst[cs-rs:]
	}
	return dst
}

func (b *base62Encoder) Encode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.encode(src), nil
}

func (b *base62Encoder) maxDecodedLen(n int) int {
	return int(math.Ceil(math.Log(62) / math.Log(256) * float64(n)))
}

func (b *base62Encoder) deocde(src []byte) ([]byte, error) {
	rs := 0
	cs := b.maxDecodedLen(len(src))
	dst := make([]byte, cs)
	for k := range src {
		v, ok := b.decodeMap[src[k]]
		if !ok {
			return nil, base.ErrEncodedText(codec, src[k], k)
		}
		c := 0
		for j := cs - 1; j >= 0 && (v != 0 || c < rs); j-- {
			v += stdEncoderSize * int(dst[j])
			dst[j] = byte(v % 256)
			v /= 256
			c++
		}
		rs = c
	}
	if cs > rs {
		return dst[cs-rs:], nil
	}
	return dst, nil
}

func (b *base62Encoder) Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.deocde(src)
}
