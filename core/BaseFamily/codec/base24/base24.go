// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/13 22:06:37                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base24                                                                                                      *
// * File: base24.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base24

import (
	"CipT/core/BaseFamily/codec/base"
	"encoding/binary"
	"github.com/pkg/errors"
)

const (
	codec          = "base24"
	stdEncoder     = "ZAC2B3EF4GH5TK67P8RS9WXY"
	stdEncoderSize = 24
)

var (
	errEncodedTextLength = errors.New("codec/base24: decode data length is not multiple of 2.")
	errEncodedText       = "codec/base24: decode data include not in encoder character: %+v, pos: %d."
)

var StdCodec, _ = NewCodec(stdEncoder)

type base24Codex struct {
	encodeMap [24]byte
	decodeMap map[byte]int
}

func NewCodec(encoder string) (base.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base.ErrEncoderSize(codec, stdEncoderSize)
	}
	b := &base24Codex{decodeMap: make(map[byte]int, stdEncoderSize)}
	mp := make(map[rune]struct{}, stdEncoderSize)
	for k, v := range encoder {
		if base.IsIllegalCharacter(v) {
			return nil, base.ErrEncoderIllegalChar(codec)
		}
		mp[v] = struct{}{}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
	}
	if len(mp) != stdEncoderSize {
		return nil, base.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base24Codex) maxEncodedLen(n int) int {
	if n%4 == 0 {
		return n / 4 * 7
	} else {
		return (n/4 + 1) * 7
	}
}

func (b *base24Codex) encode(dst, src []byte) {
	size := len(src)
	pad := 4 - size%4
	if pad == 4 {
		pad = 0
	}
	opLength := size + pad
	tmp := make([]byte, opLength)
	copy(tmp, src)
	idx := 0
	for i := 0; i < opLength; i += 4 {
		chunk := binary.BigEndian.Uint32(tmp[i : i+4])
		for j := 0; j < 7; j++ {
			idx = i/4*7 + (6 - j)
			dst[idx] = b.encodeMap[chunk%24]
			chunk /= 24
		}
	}
	return
}

func (b *base24Codex) Encode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.maxEncodedLen(len(src)))
	b.encode(dst, src)
	return dst, nil
}

func (b *base24Codex) decodedLen(n int) int {
	return n / 7 * 4
}

func (b *base24Codex) decode(dst, src []byte) (int, error) {
	elem, ok := b.decodeMap[src[0]]
	if !ok {
		return 0, errors.Errorf(errEncodedText, src[0], 0)
	}
	dst[0] = byte(elem)
	k := 0
	for i := 0; i < len(src); i += 7 {
		var chunk uint32 = 0
		for j := 0; j < 7; j++ {
			chunk *= 24
			elem, ok := b.decodeMap[src[i+j]]
			if !ok {
				return 0, errors.Errorf(errEncodedText, src[i+j], i+j)
			}
			chunk += uint32(elem)
		}
		binary.BigEndian.PutUint32(dst[k:k+4], chunk)
		k += 4
	}
	nDst := len(dst)
	n := 0
	for i := nDst - 1; i > 0; i-- {
		if dst[i] == 0x0 {
			n += 1
		} else {
			break
		}
	}
	return nDst - n, nil
}

func (b *base24Codex) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	if size%7 != 0 {
		return nil, base.ErrEncodedTextSize(codec, size, 7)
	}
	dst := make([]byte, b.decodedLen(size))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
