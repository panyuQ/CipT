// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/06 22:15:16                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base16                                                                                                      *
// * File: base16.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base16

import (
	base2 "CipT/Core/NoKey/BaseFamily/codec/base"
)

const (
	codec          = "base16"
	stdEncoder     = "0123456789abcdef"
	stdEncoderSize = 16
)

var StdCodec, _ = NewCodec(stdEncoder)

type base16Codec struct {
	encodeMap [16]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base2.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base2.ErrEncoderSize(codec, stdEncoderSize)
	}
	b := &base16Codec{decodeMap: make(map[byte]int)}
	mp := make(map[rune]struct{}, 16)
	for k, v := range encoder {
		if base2.IsIllegalCharacter(v) {
			return nil, base2.ErrEncoderIllegalChar(codec)
		}
		b.decodeMap[byte(v)] = k
		mp[v] = struct{}{}
		b.encodeMap[k] = byte(v)
	}
	if len(encoder) != len(mp) {
		return nil, base2.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base16Codec) encodedLen(n int) int { return n * 2 }

func (b *base16Codec) encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = b.encodeMap[v>>4]
		dst[j+1] = b.encodeMap[v&0x0F]
		j += 2
	}
	return len(src) * 2
}

func (b *base16Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodedLen(len(src)))
	n := b.encode(dst, src)
	return dst[:n], nil
}

func (b *base16Codec) decodedLen(n int) int { return n / 2 }

func (b *base16Codec) decode(dst, src []byte) (int, error) {
	i, j := 0, 1
	for ; j < len(src); j += 2 {
		p := src[j-1]
		q := src[j]
		a, ok := b.decodeMap[p]
		if !ok {
			return 0, base2.ErrEncodedText(codec, p, j-1)
		}
		b, ok := b.decodeMap[q]
		if !ok {
			return 0, base2.ErrEncodedText(codec, q, j)
		}
		dst[i] = byte((a << 4) | b)
		i++
	}
	return i, nil
}

func (b *base16Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size <= 0 {
		return []byte{}, nil
	}
	if size%2 != 0 {
		return nil, base2.ErrEncodedTextSize(codec, size, 2)
	}
	dst := make([]byte, b.decodedLen(len(src)))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
