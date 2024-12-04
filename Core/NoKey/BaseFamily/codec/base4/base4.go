// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/26 21:27:21                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base4                                                                                                       *
// * File: base4.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base4

import (
	base2 "CipT/Core/NoKey/BaseFamily/codec/base"
)

const (
	codec          = "base4"
	stdEncoder     = "0123"
	stdEncoderSize = 4
	stdBlock       = 4
)

var StdCodec, _ = NewCodec(stdEncoder)

type base4Codec struct {
	encodeMap [4]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base2.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base2.ErrEncoderSize(codec, stdEncoderSize)
	}

	mp := make(map[rune]struct{}, stdEncoderSize)
	b := &base4Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	for k, v := range encoder {
		if base2.IsIllegalCharacter(v) {
			return nil, base2.ErrEncoderIllegalChar(codec)
		}
		mp[v] = struct{}{}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
	}
	if len(mp) != stdEncoderSize {
		return nil, base2.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base4Codec) encodeLen(n int) int {
	return n * 4
}

func (b *base4Codec) endoce(dst, src []byte) {
	for k, v := range src {
		dst[k*4] = b.encodeMap[v>>6&0x3]
		dst[k*4+1] = b.encodeMap[v>>4&0x3]
		dst[k*4+2] = b.encodeMap[v>>2&0x3]
		dst[k*4+3] = b.encodeMap[v&0x3]
	}
}

func (b *base4Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodeLen(len(src)))
	b.endoce(dst, src)
	return dst, nil
}

func (b *base4Codec) decodeLen(n int) int {
	return n / 4
}

func (b *base4Codec) decode(dst, src []byte) (int, error) {
	nDst, val := 0, 0
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return 0, base2.ErrEncodedText(codec, v, k)
		}
		lRsh := 6 - k%stdBlock*2
		if lRsh >= 0 {
			val |= elem << lRsh
		}
		if lRsh == 0 {
			dst[nDst] = byte(val)
			nDst += 1
			val = 0
		}
	}

	return nDst, nil
}

func (b *base4Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	if size%stdBlock != 0 {
		return nil, base2.ErrEncodedTextSize(codec, size, stdEncoderSize)
	}
	dst := make([]byte, b.decodeLen(size))
	n, err := b.decode(dst, src)
	return dst[:n], err
}
