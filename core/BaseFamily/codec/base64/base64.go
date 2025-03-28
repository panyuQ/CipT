// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/25 22:46:00                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base64                                                                                                      *
// * File: base64.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base64

import (
	base2 "CipT/core/BaseFamily/codec/base"
	"encoding/binary"
	"github.com/pkg/errors"
	"strconv"
)

const (
	codec          = "base64"
	stdEncoder     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	urlEncoder     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	stdEncoderSize = 64
	errByteFormat  = "codec/base64: illegal base64 data at input byte, pos:%d."
)

var StdCodec, _ = NewCodec(stdEncoder, base2.StdPadding)
var UrlCodec, _ = NewCodec(urlEncoder, base2.StdPadding)
var StdRawCodec, _ = NewCodec(stdEncoder, base2.NotPadding)
var UrlRawCodec, _ = NewCodec(urlEncoder, base2.NotPadding)

type base64Codec struct {
	encodeMap [64]byte
	decodeMap map[byte]int
	padding   rune
	strict    bool
}

func NewCodec(encoder string, padding rune) (base2.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base2.ErrEncoderSize(codec, stdEncoderSize)
	}

	if base2.IsIllegalCharacter(padding) {
		return nil, base2.ErrPaddingIllegalChar(codec)
	}

	b := &base64Codec{decodeMap: make(map[byte]int, stdEncoderSize), padding: padding}
	mp := make(map[rune]struct{}, stdEncoderSize)

	for k, v := range encoder {
		if base2.IsIllegalCharacter(v) {
			return nil, base2.ErrEncoderIllegalChar(codec)
		}
		b.encodeMap[k] = byte(v)
		b.decodeMap[byte(v)] = k
		mp[v] = struct{}{}
	}
	if len(mp) != stdEncoderSize {
		return nil, base2.ErrEncoderRepeatChar(codec)
	}
	return b, nil
}

func (b *base64Codec) encodedLen(n int) int {
	if b.padding == base2.NotPadding {
		return (n*8 + 5) / 6 // minimum # chars at 6 bits per char
	}
	return (n + 2) / 3 * 4 // minimum # 4-char quanta, 3 bytes each
}

func (b *base64Codec) encode(dst, src []byte) {
	di, si := 0, 0
	n := (len(src) / 3) * 3
	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])

		dst[di+0] = b.encodeMap[val>>18&0x3F]
		dst[di+1] = b.encodeMap[val>>12&0x3F]
		dst[di+2] = b.encodeMap[val>>6&0x3F]
		dst[di+3] = b.encodeMap[val&0x3F]

		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return
	}
	// Add the remaining small block
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = b.encodeMap[val>>18&0x3F]
	dst[di+1] = b.encodeMap[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = b.encodeMap[val>>6&0x3F]
		if b.padding != base2.NotPadding {
			dst[di+3] = byte(b.padding)
		}
	case 1:
		if b.padding != base2.NotPadding {
			dst[di+2] = byte(b.padding)
			dst[di+3] = byte(b.padding)
		}
	}
}

func (b *base64Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodedLen(len(src)))
	b.encode(dst, src)
	return dst, nil
}

func (b *base64Codec) decodedLen(n int) int {
	if b.padding == base2.NotPadding {
		// Unpadded data may end with partial block of 2-3 characters.
		return n * 6 / 8
	}
	// Padded base64 should always be a multiple of 4 characters in length.
	return n / 4 * 3
}

func (b *base64Codec) decodeQuantum(dst, src []byte, si int) (nsi, n int, err error) {
	var dbuf [4]byte
	dlen := 4

	for j := 0; j < len(dbuf); j++ {
		if len(src) == si {
			switch {
			case j == 0:
				return si, 0, nil
			case j == 1, b.padding != base2.NotPadding:
				return si, 0, errors.Errorf(errByteFormat, si-j)
			}
			dlen = j
			break
		}
		in := src[si]
		si++
		out, ok := b.decodeMap[in]
		if ok {
			dbuf[j] = byte(out)
			continue
		}

		if rune(in) != b.padding {
			return si, 0, errors.Errorf(errByteFormat, si-1)
		}
		switch j {
		case 0, 1:
			return si, 0, errors.Errorf(errByteFormat, si-1)
		case 2:
			if si == len(src) {
				// not enough padding
				return si, 0, errors.Errorf(errByteFormat, len(src))
			}
			if rune(src[si]) != b.padding {
				// incorrect padding
				return si, 0, errors.Errorf(errByteFormat, si-1) //CorruptInputError(si - 1)
			}

			si++
		}
		if si < len(src) {
			// trailing garbage
			err = errors.Errorf(errByteFormat, si)
		}
		dlen = j
		break
	}
	// Convert 4x 6bit source bytes into 3 bytes
	val := uint(dbuf[0])<<18 | uint(dbuf[1])<<12 | uint(dbuf[2])<<6 | uint(dbuf[3])
	dbuf[2], dbuf[1], dbuf[0] = byte(val>>0), byte(val>>8), byte(val>>16)
	switch dlen {
	case 4:
		dst[2] = dbuf[2]
		dbuf[2] = 0
		fallthrough
	case 3:
		dst[1] = dbuf[1]
		if b.strict && dbuf[2] != 0 {
			return si, 0, errors.Errorf(errByteFormat, si-1)
		}
		dbuf[1] = 0
		fallthrough
	case 2:
		dst[0] = dbuf[0]
		if b.strict && (dbuf[1] != 0 || dbuf[2] != 0) {
			return si, 0, errors.Errorf(errByteFormat, si-2)
		}
	}

	return si, dlen - 1, err
}

func (b *base64Codec) assemble32(src []byte) (uint32, bool) {
	var val uint32
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return 0, false
		}
		val |= uint32(elem) << (26 - k%4*6)
	}
	return val, true
}

func (b *base64Codec) assemble64(src []byte) (uint64, bool) {
	var val uint64
	for k, v := range src {
		elem, ok := b.decodeMap[v]
		if !ok {
			return 0, false
		}
		val |= uint64(elem) << (58 - k%8*6)
	}
	return val, true
}

func (b *base64Codec) decode(dst, src []byte) (n int, err error) {
	si := 0
	for strconv.IntSize >= 64 && len(src)-si >= 8 && len(dst)-n >= 8 {
		tmp := src[si : si+8]
		if dn, ok := b.assemble64(tmp); ok {
			binary.BigEndian.PutUint64(dst[n:], dn)
			n += 6
			si += 8
		} else {
			var ninc int
			si, ninc, err = b.decodeQuantum(dst[n:], src, si)
			n += ninc
			if err != nil {
				return n, err
			}
		}
	}
	for len(src)-si >= 4 && len(dst)-n >= 4 {
		tmp := src[si : si+4]
		if dn, ok := b.assemble32(tmp); ok {
			binary.BigEndian.PutUint32(dst[n:], dn)
			n += 3
			si += 4
		} else {
			var ninc int
			si, ninc, err = b.decodeQuantum(dst[n:], src, si)
			n += ninc
			if err != nil {
				return n, err
			}
		}
	}

	for si < len(src) {
		var ninc int
		si, ninc, err = b.decodeQuantum(dst[n:], src, si)
		n += ninc
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func (b *base64Codec) Decode(src []byte) ([]byte, error) {
	size := len(src)
	if size == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.decodedLen(size))
	n, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
