// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 21:43:14                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base2                                                                                                       *
// * File: base2.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base2

import (
	"CipT/codec/base"
)

const (
	codec          = "base2"
	stdEncoder     = "01"
	stdEncoderSize = 2
	stdBlockLen    = 8
)

var StdCodec, _ = NewCodec(stdEncoder)

type base2Codec struct {
	encodeMap [2]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}
	if encoder[0] == encoder[1] {
		return nil, base.ErrEncoderRepeatChar(codec)
	}
	b := &base2Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
		}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
	}
	return b, nil
}

func (b *base2Codec) encodeLen(n int) int { return n * 8 }

func (b *base2Codec) encode(dst, src []byte) {
	for k, v := range src {
		dst[k*8] = b.encodeMap[v>>7&0x1]
		dst[k*8+1] = b.encodeMap[v>>6&0x1]
		dst[k*8+2] = b.encodeMap[v>>5&0x1]
		dst[k*8+3] = b.encodeMap[v>>4&0x1]
		dst[k*8+4] = b.encodeMap[v>>3&0x1]
		dst[k*8+5] = b.encodeMap[v>>2&0x1]
		dst[k*8+6] = b.encodeMap[v>>1&0x1]
		dst[k*8+7] = b.encodeMap[v>>0&0x1]
	}
}

func (b *base2Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodeLen(len(src)))
	b.encode(dst, src)
	return dst, nil
}

func (b *base2Codec) decodeLen(n int) int { return n / 8 }

func (b *base2Codec) decode(dst, src []byte) (int, error) {
	nDst, val := 0, 0
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return 0, base.ErrEncodedText(codec, v, k)
		}

		lRsh := 7 - k%stdBlockLen
		if lRsh >= 0 {
			val |= elem << lRsh
		}

		if lRsh == 0 {
			dst[nDst] = byte(val)
			nDst += 1
			val = 0
		}
	}
	return 0, nil
}

func (b *base2Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	if size%stdBlockLen != 0 {
		return nil, base.ErrEncodedTextSize(codec, size, stdBlockLen)
	}
	dst := make([]byte, b.decodeLen(size))
	n, err := b.decode(dst, src)
	return dst[:n], err
}
