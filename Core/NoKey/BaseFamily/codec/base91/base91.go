// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 21:36:23                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base91                                                                                                      *
// * File: base91.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base91

import (
	"CipT/Core/NoKey/BaseFamily/codec/base"
	"CipT/Core/NoKey/BaseFamily/codec/internal/errs"
	"CipT/Core/NoKey/BaseFamily/codec/internal/util"
	"math"
)

const (
	codec          = "base91"
	stdEncoder     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""
	stdEncoderSize = 91
)

var StdCodec, _ = NewCodec(stdEncoder)

type base91Codec struct {
	encodeMap [91]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, errs.ErrEncoderSize(codec, stdEncoderSize)
	}
	mp := make(map[rune]struct{}, stdEncoderSize)
	b := &base91Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	for k, v := range encoder {
		if util.IsIllegalCharacter(v) {
			return nil, errs.ErrEncoderIllegalChar(codec)
		}
		mp[v] = struct{}{}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
	}
	if len(mp) != stdEncoderSize {
		return nil, errs.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base91Codec) maxEncodedLen(n int) int {
	return int(math.Ceil(float64(n) * 16.0 / 13.0))
}

func (b *base91Codec) encode(dst, src []byte) (int, error) {
	var queue, numBits uint
	n := 0
	for i := 0; i < len(src); i++ {
		queue |= uint(src[i]) << numBits
		numBits += 8
		if numBits > 13 {
			v := queue & 8191
			if v > 88 {
				queue >>= 13
				numBits -= 13
			} else {
				v = queue & 16383
				queue >>= 14
				numBits -= 14
			}
			dst[n] = b.encodeMap[v%91]
			dst[n+1] = b.encodeMap[v/91]
			n += 2
		}
	}
	if numBits > 0 {
		dst[n] = b.encodeMap[queue%91]
		n++
		if numBits > 7 || queue > 90 {
			dst[n] = b.encodeMap[queue/91]
			n++
		}
	}
	return n, nil
}

func (b *base91Codec) Encode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.maxEncodedLen(size))
	n, err := b.encode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

func (b *base91Codec) maxDecodedLen(n int) int {
	return int(math.Ceil(float64(n) * 14.0 / 16.0))
}

func (b *base91Codec) decode(dst, src []byte) (int, error) {
	var queue, numBits uint
	v, n := -1, 0
	for i := 0; i < len(src); i++ {
		elem, ok := b.decodeMap[src[i]]
		if !ok {
			return 0, errs.ErrEncodedText(codec, src[i], i)
		}
		if v == -1 {
			v = elem
		} else {
			v += elem * 91
			queue |= uint(v) << numBits
			if (v & 8191) > 88 {
				numBits += 13
			} else {
				numBits += 14
			}

			for ok := true; ok; ok = numBits > 7 {
				dst[n] = byte(queue)
				n++
				queue >>= 8
				numBits -= 8
			}
			v = -1
		}
	}
	if v != -1 {
		dst[n] = byte(queue | uint(v)<<numBits)
		n++
	}
	return n, nil
}

func (b *base91Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.maxEncodedLen(size))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
