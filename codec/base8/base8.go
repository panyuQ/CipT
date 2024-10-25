// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/22 22:44:32                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base8                                                                                                       *
// * File: base8.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base8

import (
	"CipT/codec/base"
	"github.com/pkg/errors"
)

const (
	StdEncoder     = "01234567"
	stdEncoderSize = 8
	codec          = "base8"
)

var (
	errEncodedDataSize = errors.New("codec/base8: invalid encoded data length.")
)

var StdCodec, _ = NewCodec(StdEncoder, base.StdPadding)

type base8Codec struct {
	encodeMap [8]byte
	decodeMap map[byte]int
	padding   rune
}

func NewCodec(encoder string, padding rune) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}

	if base.IsIllegalCharacter(padding) {
		return nil, base.ErrPaddingIllegalChar(codec)
	}

	mp := make(map[rune]struct{}, stdEncoderSize)
	b := &base8Codec{decodeMap: make(map[byte]int, stdEncoderSize), padding: padding}
	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
		}
		if v == padding {
			return nil, base.ErrPaddingAlphabet(codec)
		}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
		mp[v] = struct{}{}
	}

	if len(mp) != stdEncoderSize {
		return nil, base.ErrEncoderRepeatChar(codec)
	}

	return b, nil
}

func (b *base8Codec) maxEncodeLen(n int) int {
	if n%3 == 0 {
		return n / 3 * 8
	} else {
		return (n/3 + 1) * 8
	}
}

func (b *base8Codec) encode(dst, src []byte) int {
	length := len(src)
	if length <= 0 {
		return 0
	}

	nDst := 0
	// 每3字节一组的部分
	opLength := length / 3
	{
		for i := 0; i < opLength; i++ {
			val := (uint(src[i*3]) << 16) | (uint(src[i*3+1]) << 8) | (uint(src[i*3+2]))
			dst[i*8] = b.encodeMap[(val>>21)&0x07]
			dst[i*8+1] = b.encodeMap[(val>>18)&0x07]
			dst[i*8+2] = b.encodeMap[(val>>15)&0x07]
			dst[i*8+3] = b.encodeMap[(val>>12)&0x07]
			dst[i*8+4] = b.encodeMap[(val>>9)&0x07]
			dst[i*8+5] = b.encodeMap[(val>>6)&0x07]
			dst[i*8+6] = b.encodeMap[(val>>3)&0x07]
			dst[i*8+7] = b.encodeMap[(val)&0x07]
			nDst += 8
		}
	}
	// 剩余部分
	{
		remain := length - (opLength * 3)
		if remain == 0 {
			return nDst
		}
		val := uint(src[opLength*3]) << 16
		if remain == 2 {
			val |= uint(src[opLength*3+1]) << 8
		}
		dst[nDst] = b.encodeMap[val>>21&0x7]
		dst[nDst+1] = b.encodeMap[val>>18&0x7]
		dst[nDst+2] = b.encodeMap[val>>15&0x7]
		switch remain {
		case 2:
			dst[nDst+3] = b.encodeMap[val>>12&0x7]
			dst[nDst+4] = b.encodeMap[val>>9&0x7]
			dst[nDst+5] = b.encodeMap[val>>6&0x7]
			if b.padding == base.NotPadding {
				nDst += 6
				return nDst
			}
			dst[nDst+6] = byte(b.padding)
			dst[nDst+7] = byte(b.padding)
			nDst += 8
			return nDst
		case 1:
			if b.padding == base.NotPadding {
				nDst += 3
				return nDst
			}
			dst[nDst+3] = byte(b.padding)
			dst[nDst+4] = byte(b.padding)
			dst[nDst+5] = byte(b.padding)
			dst[nDst+6] = byte(b.padding)
			dst[nDst+7] = byte(b.padding)
			nDst += 8
			return nDst
		}
	}

	return nDst
}

func (b *base8Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.maxEncodeLen(len(src)))
	n := b.encode(dst, src)
	return dst[:n], nil
}

func (b *base8Codec) maxDecodeLen(n int) int {
	if n%8 == 0 {
		return n / 8 * 3
	} else {
		return (n/8 + 1) * 3
	}
}

func (b *base8Codec) decode(dst, src []byte) (int, error) {
	length := len(src)
	pad := 0
	for i := length - 1; i >= 0; i-- {
		if src[i] == byte(b.padding) {
			pad++
		} else {
			break
		}
	}

	{
		remain := (length - pad) % 8
		if remain != 0 && remain != 3 && remain != 6 {
			return 0, base.ErrDecodeSrcDataSize(codec, length)
		}
	}

	opLength := ((length - pad) / 8) * 8

	nDst := 0
	{
		val := 0
		for i := 0; i < opLength; i++ {
			v, ok := b.decodeMap[src[i]]
			if !ok {
				return 0, base.ErrEncodedText(codec, src[i], i)
			}
			lRsh := 21 - i%8*3
			if lRsh >= 0 {
				val |= v << lRsh
			}
			if lRsh == 0 {
				dst[nDst] = byte(val >> 16)
				dst[nDst+1] = byte(val >> 8)
				dst[nDst+2] = byte(val)
				nDst += 3
				val = 0
			}
		}
	}
	{
		remain := length - pad - opLength
		if remain == 0 {
			return nDst, nil
		}
		if remain != 3 && remain != 6 {
			return 0, base.ErrDecodeSrcDataSize(codec, length)
		}
		val, err := b.assemble(0, src, opLength, opLength+1, opLength+2)
		if err != nil {
			return 0, err
		}
		switch remain {
		case 6:
			val, err = b.assemble(val, src, opLength+3, opLength+4, opLength+5)
			if err != nil {
				return 0, err
			}
			dst[nDst] = byte(val >> 16)
			dst[nDst+1] = byte(val >> 8)
			nDst += 2
			return nDst, nil
		case 3:
			dst[nDst] = byte(val >> 16)
			nDst += 1
			return nDst, nil
		}
	}
	return nDst, nil
}

func (b *base8Codec) assemble(val int, src []byte, idx ...int) (int, error) {
	for _, v := range idx {
		elem, ok := b.decodeMap[src[v]]
		if !ok {
			return 0, base.ErrEncodedText(codec, src[v], v)
		}
		lRsh := 21 - (v % 8 * 3)
		val |= elem << lRsh
	}
	return val, nil
}

func (b *base8Codec) Decode(src []byte) ([]byte, error) {
	srcLen := len(src)
	if srcLen <= 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.maxDecodeLen(srcLen))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
