// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/12 22:29:48                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base58                                                                                                      *
// * File: base58.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base58

import (
	"CipT/codec/base"
	"math"
	"math/big"
)

const (
	codec          = "base58"
	stdEncoder     = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	stdEncoderSize = 58
)

var StdCodec, _ = NewCodec(stdEncoder)

type bas58Codec struct {
	encodedMap [stdEncoderSize]byte
	decodedMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}

	b := &bas58Codec{decodedMap: make(map[byte]int, stdEncoderSize)}
	mp := make(map[byte]struct{}, stdEncoderSize)

	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
		}
		b.encodedMap[k] = byte(v)
		b.decodedMap[byte(v)] = k
		mp[byte(v)] = struct{}{}
	}

	if len(mp) != stdEncoderSize {
		return nil, base.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *bas58Codec) maxEncodedLen(n int) int {
	return int(math.Ceil(math.Log(256) / math.Log(stdEncoderSize) * float64(n)))
}

func (b *bas58Codec) encode(src []byte) []byte {
	srcInt := big.NewInt(0).SetBytes(src)
	mod, radix := big.NewInt(0), big.NewInt(stdEncoderSize)
	zero := big.NewInt(0)
	var dst = make([]byte, 0, b.maxEncodedLen(len(src)))
	for srcInt.Cmp(zero) > 0 {
		srcInt.DivMod(srcInt, radix, mod)
		dst = append(dst, b.encodedMap[mod.Int64()])
	}
	for _, v := range src {
		if v != 0 {
			break
		}
		dst = append(dst, b.encodedMap[0])
	}
	nDst := len(dst)
	for i := 0; i < nDst/2; i++ {
		dst[i], dst[nDst-1-i] = dst[nDst-1-i], dst[i]
	}
	return dst
}

func (b *bas58Codec) Encode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.encode(src), nil
}

func (b *bas58Codec) maxDecodedLen(n int) int {
	return int(math.Ceil(math.Log(stdEncoderSize) / math.Log(256) * float64(n)))
}

func (b *bas58Codec) decode(src []byte) ([]byte, error) {
	bigInt := big.NewInt(0)
	radix := big.NewInt(stdEncoderSize)
	for k, v := range src {
		elem, ok := b.decodedMap[v]
		if !ok {
			return nil, base.ErrEncodedText(codec, v, k)
		}
		bigInt.Mul(bigInt, radix)
		bigInt.Add(bigInt, big.NewInt(int64(elem)))
	}
	tmpBytes := bigInt.Bytes()
	var numZeros int
	for numZeros = 0; numZeros < len(src); numZeros++ {
		if src[numZeros] != b.encodedMap[0] {
			break
		}
	}
	length := numZeros + len(tmpBytes)
	dst := make([]byte, length)
	copy(dst[numZeros:], tmpBytes)
	return dst, nil
}

func (b *bas58Codec) Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.decode(src)
}
