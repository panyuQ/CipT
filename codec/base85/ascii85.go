// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/20 18:52:22                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: ascii85                                                                                                     *
// * File: ascii855.go                                                                                                 *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base85

import (
	"CipT/codec/base"
	"github.com/pkg/errors"
)

const (
	codec          = "base85"
	StdEncoder     = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstu"
	stdEncoderSize = 85
)

type ascii85Codec struct {
	encodeMap [85]byte
	deocdeMap [256]byte
}

var (
	errEncoderSize            = errors.New("codec/ascii85: encoding alphabet is not 85-bytes long.")
	errEncoderRepeatCharacter = errors.New("codec/ascii85: encoding alphabet has repeat character.")
	errEncoderCharacter       = errors.New("codec/ascii85: encoding alphabet contains illegal character.")
	decodeBase                = [5]uint32{85 * 85 * 85 * 85, 85 * 85 * 85, 85 * 85, 85, 1}
	invalidInputFormat        = "codec/ascii85: illegal ascii85 data at input byte, pos:%d."
	StdCodec, _               = NewCodec(StdEncoder)
)

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}

	b := new(ascii85Codec)
	for k, v := range encoder {

		if base.IsIllegalCharacter(v) {
			return nil, errEncoderCharacter
		}
		b.deocdeMap[v] = byte(k)
	}
	copy(b.encodeMap[:], encoder)
	return b, nil
}

func (b *ascii85Codec) encodedLen(n int) int {
	s := n / 4
	r := n % 4
	if r > 0 {
		return s*5 + 5 - (4 - r)
	} else {
		return s * 5
	}
}

func (b *ascii85Codec) encodeChunk(dst, src []byte) int {
	if len(src) == 0 {
		return 0
	}
	var val uint32
	switch len(src) {
	default:
		val |= uint32(src[3])
		fallthrough
	case 3:
		val |= uint32(src[2]) << 8
		fallthrough
	case 2:
		val |= uint32(src[1]) << 16
		fallthrough
	case 1:
		val |= uint32(src[0]) << 24
	}
	buf := [5]byte{0, 0, 0, 0, 0}
	for i := 4; i >= 0; i-- {
		r := val % 85
		val /= 85
		buf[i] = b.encodeMap[r]
	}
	m := b.encodedLen(len(src))
	copy(dst[:], buf[:m])
	return m
}

func (b *ascii85Codec) encode(dst, src []byte) int {
	n := 0
	for len(src) > 0 {
		if len(src) < 4 {
			n += b.encodeChunk(dst, src)
			return n
		}
		n += b.encodeChunk(dst[:5], src[:4])
		src = src[4:]
		dst = dst[5:]
	}
	return n
}

func (b *ascii85Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodedLen(len(src)))
	b.encode(dst, src)
	return dst, nil
}

func (b *ascii85Codec) decodedLen(n int) int {
	s := n / 5
	r := n % 5
	if r > 0 {
		return s*4 + 4 - (5 - r)
	} else {
		return s * 4
	}
}

func (b *ascii85Codec) decodeChunk(dst, src []byte) (int, int) {
	if len(src) == 0 {
		return 0, 0
	}
	var val uint32
	m := b.decodedLen(len(src))
	buf := [5]byte{84, 84, 84, 84, 84}
	for i := 0; i < len(src); i++ {
		e := b.deocdeMap[src[i]]
		if e == 0xFF {
			return 0, i + 1
		}
		buf[i] = e
	}
	for i := 0; i < 5; i++ {
		r := buf[i]
		val += uint32(r) * decodeBase[i]
	}
	switch m {
	default:
		dst[3] = byte(val & 0xFF)
		fallthrough
	case 3:
		dst[2] = byte((val >> 8) & 0xFF)
		fallthrough
	case 2:
		dst[1] = byte((val >> 16) & 0xFF)
		fallthrough
	case 1:
		dst[0] = byte((val >> 24) & 0xFF)
	}
	return m, 0
}

func (b *ascii85Codec) decode(dst, src []byte) (int, error) {
	f := 0
	t := 0
	for len(src) > 0 {
		if len(src) < 5 {
			w, n := b.decodeChunk(dst, src)
			if n > 0 {
				return t, errors.Errorf(invalidInputFormat, n+f)
			}
			return t + w, nil
		}
		_, n := b.decodeChunk(dst[:4], src[:5])
		if n > 0 {
			return t, errors.Errorf(invalidInputFormat, n+f)
		} else {
			t += 4
			f += 5
			src = src[5:]
			dst = dst[4:]
		}
	}
	return t, nil
}

func (b *ascii85Codec) Decode(src []byte) ([]byte, error) {
	dst := make([]byte, b.decodedLen(len(src)))
	_, err := b.decode(dst, src)
	return dst, err
}
