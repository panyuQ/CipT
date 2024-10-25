// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/25 20:59:55                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base36                                                                                                      *
// * File: base36.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base36

import (
	"CipT/codec/base"
	"math/big"
)

const (
	codec          = "base36"
	stdEncoder     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	stdEncoderSize = 36
)

var StdCodec, _ = NewCodec(stdEncoder)

type base36Codec struct {
	encodeMap [stdEncoderSize]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}
	mp := make(map[rune]struct{}, stdEncoderSize)
	b := &base36Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
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

func (b *base36Codec) encodedLen(n int) int {
	return n * 136 / 100
}

func (b *base36Codec) encode(src []byte) []byte {
	x := new(big.Int)
	x.SetBytes(src)
	dst := make([]byte, 0, b.encodedLen(len(src)))
	mod, radix := big.NewInt(0), big.NewInt(stdEncoderSize)
	zero := big.NewInt(0)
	for x.Cmp(zero) > 0 {
		x.DivMod(x, radix, mod)
		dst = append(dst, b.encodeMap[mod.Int64()])
	}

	for _, i := range src {
		if i != 0 {
			break
		}
		dst = append(dst, b.encodeMap[0])
	}

	nDst := len(dst)
	for i := 0; i < nDst/2; i++ {
		dst[i], dst[nDst-1-i] = dst[nDst-1-i], dst[i]
	}
	return dst
}

func (b *base36Codec) Encode(src []byte) ([]byte, error) {
	if len(src) <= 0 {
		return []byte{}, nil
	}
	return b.encode(src), nil
}

func (b *base36Codec) decode(src []byte) ([]byte, error) {
	bigInt := big.NewInt(0)
	radix := big.NewInt(stdEncoderSize)
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return nil, base.ErrEncodedText(codec, v, k)
		}
		bigInt.Mul(bigInt, radix)
		bigInt.Add(bigInt, big.NewInt(int64(elem)))
	}
	tmpBytes := bigInt.Bytes()
	var numZeros int
	for numZeros = 0; numZeros < len(src); numZeros++ {
		if src[numZeros] != b.encodeMap[0] {
			break
		}
	}
	length := numZeros + len(tmpBytes)
	dst := make([]byte, length)
	copy(dst[numZeros:], tmpBytes)
	return dst, nil
}

func (b *base36Codec) Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.decode(src)
}
