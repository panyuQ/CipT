// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/03 23:40:52                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: internal                                                                                                    *
// * File: basex.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package basex

import (
	"CipT/codec/base"
	"math"
	"math/big"
)

type Basex struct {
	name        string
	radix       *big.Int
	encoderSize int
	encodeMap   []byte
	decodeMap   map[byte]int
}

func NewBasex(name, encoder string, encodeSize int) (*Basex, error) {
	if len(encoder) != encodeSize {
		return nil, base.ErrEncoderSize(name, encodeSize)
	}

	b := &Basex{
		name:        name,
		radix:       big.NewInt(int64(encodeSize)),
		encoderSize: encodeSize,
		encodeMap:   make([]byte, encodeSize),
		decodeMap:   make(map[byte]int, encodeSize)}
	mp := make(map[byte]struct{}, encodeSize)

	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(name)
		}
		b.encodeMap[k] = byte(v)
		b.decodeMap[byte(v)] = k
		mp[byte(v)] = struct{}{}
	}

	if len(mp) != encodeSize {
		return nil, base.ErrEncoderRepeatChar(name)
	}
	return b, nil
}

func (b *Basex) maxEncodedLen(n int) int {
	return int(math.Ceil(math.Log(256) / math.Log(float64(b.encoderSize)) * float64(n)))
}

func (b *Basex) Encode(src []byte) []byte {
	srcInt := big.NewInt(0).SetBytes(src)
	mod := big.NewInt(0)
	zero := big.NewInt(0)
	var dst = make([]byte, 0, b.maxEncodedLen(len(src)))
	for srcInt.Cmp(zero) > 0 {
		srcInt.DivMod(srcInt, b.radix, mod)
		dst = append(dst, b.encodeMap[mod.Int64()])
	}
	for _, v := range src {
		if v != 0 {
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

func (b *Basex) maxDecodedLen(n int) int {
	return int(math.Ceil(math.Log(float64(b.encoderSize)) / math.Log(256) * float64(n)))
}

func (b *Basex) Decode(src []byte) ([]byte, error) {
	bigInt := big.NewInt(0)
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return nil, base.ErrEncodedText(b.name, v, k)
		}
		bigInt.Mul(bigInt, b.radix)
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
