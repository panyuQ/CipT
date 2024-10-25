// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 22:11:48                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base92                                                                                                      *
// * File: base92.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base92

import (
	"CipT/codec/base"
	"CipT/codec/internal/errs"
	"CipT/codec/internal/util"
	"math"
	"math/big"
)

const (
	codec          = "base92"
	stdEncoder     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#|;,_~`\""
	stdEncoderSize = 92
)

var (
	StdCodec, _ = NewCodec(stdEncoder)
)

type base92Codec struct {
	encodeMap [stdEncoderSize]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, errs.ErrEncoderSize(codec, stdEncoderSize)
	}
	b := &base92Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	mp := make(map[rune]struct{}, stdEncoderSize)

	for k, v := range encoder {
		if util.IsIllegalCharacter(v) {
			return nil, errs.ErrEncoderIllegalChar(codec)
		}
		b.encodeMap[k] = byte(v)
		b.decodeMap[byte(v)] = k
		mp[v] = struct{}{}
	}

	if len(mp) != stdEncoderSize {
		return nil, errs.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base92Codec) maxEncodedLen(n int) int {
	return int(math.Ceil(math.Log(256) / math.Log(stdEncoderSize) * float64(n)))
}

func (b *base92Codec) Encode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, 0, b.maxEncodedLen(len(src)))
	x := new(big.Int).SetBytes(src)
	bigZero := big.NewInt(0)
	bigRadix := big.NewInt(stdEncoderSize)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
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
	return dst, nil
}

func (b *base92Codec) maxDecodedLen(n int) int {
	return int(math.Ceil(math.Log(stdEncoderSize) / math.Log(256) * float64(n)))
}

func (b *base92Codec) Decode(src []byte) ([]byte, error) {
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
