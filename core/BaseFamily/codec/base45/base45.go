// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/29 23:27:15                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base45                                                                                                      *
// * File: base45.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base45

import (
	"CipT/core/BaseFamily/codec/base"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
)

const (
	codec          = "base45"
	stdEncoder     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ&$%*+-./:"
	stdEncoderSize = 45
	baseX          = 45
	baseSquare     = 45 * 45
)

var StdCodec, _ = NewCodec(stdEncoder)

type base45Codec struct {
	encodeMap [45]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}
	b := &base45Codec{decodeMap: make(map[byte]int)}
	mp := make(map[rune]struct{}, 45)

	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrPaddingIllegalChar(codec)
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

func (b *base45Codec) encodePairs(in []byte) [][]byte {
	size := len(in)
	ret := make([][]byte, 0, size/2)
	for i := 0; i < size; i += 2 {
		var high, low byte
		if i+1 < size {
			high = in[i]
			low = in[i+1]
		} else {
			low = in[i]
		}
		ret = append(ret, []byte{high, low})
	}
	return ret
}

func (b *base45Codec) encodeBase45(in []byte) []byte {
	n := binary.BigEndian.Uint16(in)
	c := n % baseX
	e := (n - c) / baseSquare
	d := (n - (c + (e * baseSquare))) / baseX
	return []byte{byte(c), byte(d), byte(e)}
}

func (b *base45Codec) encodedLen(n int) int {
	if n%2 == 0 {
		return n / 2 * 3
	} else {
		return (n/2 + 1) * 3
	}
}

func (b *base45Codec) encode(dst, src []byte) int {
	pairs := b.encodePairs(src)
	size := len(pairs)
	idx := 0
	for k, pair := range pairs {
		res := b.encodeBase45(pair)
		if k+1 == size && res[2] == 0 {
			for _, r := range res[:2] {
				dst[idx] = b.encodeMap[r]
				idx++
			}
		} else {
			for _, r := range res {
				dst[idx] = b.encodeMap[r]
				idx++
			}
		}
	}
	return idx
}

func (b *base45Codec) Encode(src []byte) ([]byte, error) {
	if len(src) <= 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.encodedLen(len(src)))
	n := b.encode(dst, src)
	return dst[:n], nil
}

func (b *base45Codec) decodeChunk(in []byte) [][]byte {
	size := len(in)
	ret := make([][]byte, 0, size/2)
	for i := 0; i < size; i += 3 {
		if i+2 < size {
			ret = append(ret, []byte{in[i], in[i+1], in[i+2]})
		} else {
			ret = append(ret, []byte{in[i], in[i+1]})
		}
	}
	return ret
}

func (b *base45Codec) decodeTriplets(in [][]byte) ([]uint16, error) {
	size := len(in)
	ret := make([]uint16, 0, size)
	for pos, chunk := range in {
		if len(chunk) == 3 {
			n := int(chunk[0]) + (int(chunk[1]) * baseX) + (int(chunk[2]) * baseSquare)
			if n > math.MaxUint16 {
				return nil, errors.Errorf("codec/base45: illegal base45 data at byte pos: %d.", pos)
			}
			ret = append(ret, uint16(n))
		}
		if len(chunk) == 2 {
			n := uint16(chunk[0]) + uint16(chunk[1])*baseX
			ret = append(ret, n)
		}
	}
	return ret, nil
}

func (b *base45Codec) decodedLen(n int) int {
	if n%3 == 0 {
		return n / 3 * 2
	} else {
		return (n/3 + 1) * 2
	}
}

func (b *base45Codec) decode(dst, src []byte) (int, error) {
	size := len(src)
	bytes := make([]byte, 0, size)
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return 0, base.ErrEncodedText(codec, v, k)
		}
		bytes = append(bytes, byte(elem))
	}

	chunks := b.decodeChunk(bytes)
	triplets, err := b.decodeTriplets(chunks)
	if err != nil {
		return 0, err
	}

	tripletsSize := len(triplets)
	nDst := 0
	for i := 0; i < tripletsSize-1; i++ {
		bytes := base.Uint16ToBytes(triplets[i])
		dst[nDst] = bytes[0]
		dst[nDst+1] = bytes[1]
		nDst += 2
	}
	if size%3 == 2 {
		bytes := base.Uint16ToBytes(triplets[tripletsSize-1])
		dst[nDst] = bytes[1]
		nDst += 1
	} else {
		bytes := base.Uint16ToBytes(triplets[tripletsSize-1])
		dst[nDst] = bytes[0]
		dst[nDst+1] = bytes[1]
		nDst += 2
	}
	return nDst, nil
}

func (b *base45Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	mod := size % 3
	if mod != 0 && mod != 2 {
		return nil, errEncodedSizeMod(size, mod)
	}
	dst := make([]byte, b.decodedLen(size))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}

func errEncodedSizeMod(size, mod int) error {
	return errors.Errorf("codec/base45: invalid decodd src size n = %d, it should be n mod 3 = [0, 2] not n mod 3 = %d.", size, mod)
}
