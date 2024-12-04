// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/16 23:40:15                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base32                                                                                                      *
// * File: base32.go                                                                                                   *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base32

import (
	base2 "CipT/Core/NoKey/BaseFamily/codec/base"
)

const (
	codec      = "base32"
	stdEncoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

	stdEncoderSize = 32
	hexEncoder     = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
)

var (
	StdCodec, _    = NewCodec(stdEncoder, base2.StdPadding)
	HexCodec, _    = NewCodec(hexEncoder, base2.StdPadding)
	RawStdCodec, _ = NewCodec(stdEncoder, base2.NotPadding)
	RawHexCodec, _ = NewCodec(hexEncoder, base2.NotPadding)
)

type base32Codec struct {
	encodeMap [32]byte
	decodeMap map[byte]int
	padding   rune
}

func NewCodec(encoder string, padding rune) (base2.IEncoding, error) {
	if len(encoder) != stdEncoderSize {
		return nil, base2.ErrEncoderSize(codec, stdEncoderSize)
	}
	if base2.IsIllegalCharacter(padding) {
		return nil, base2.ErrPaddingIllegalChar(codec)
	}
	b := &base32Codec{decodeMap: make(map[byte]int, stdEncoderSize)}
	mp := make(map[rune]struct{}, stdEncoderSize)
	for k, v := range encoder {
		if base2.IsIllegalCharacter(v) {
			return nil, base2.ErrEncoderIllegalChar(codec)
		}
		if v == padding {
			return nil, base2.ErrPaddingAlphabet(codec)
		}
		mp[v] = struct{}{}
		b.decodeMap[byte(v)] = k
		b.encodeMap[k] = byte(v)
	}
	if len(mp) != stdEncoderSize {
		return nil, base2.ErrEncoderRepeatChar(codec)
	}
	b.padding = padding
	return b, nil
}

func (b *base32Codec) encodedLen(n int) int {
	if b.padding == base2.NotPadding {
		return (n*8 + 4) / 5
	}
	return (n + 4) / 5 * 8
}

func (b *base32Codec) encode(dst, src []byte) {
	for len(src) > 0 {
		var buf [8]byte

		// Unpack 8x 5-bit source blocks into a 5 byte
		// destination quantum
		switch len(src) {
		default:
			buf[7] = src[4] & 0x1F
			buf[6] = src[4] >> 5
			fallthrough
		case 4:
			buf[6] |= (src[3] << 3) & 0x1F
			buf[5] = (src[3] >> 2) & 0x1F
			buf[4] = src[3] >> 7
			fallthrough
		case 3:
			buf[4] |= (src[2] << 1) & 0x1F
			buf[3] = (src[2] >> 4) & 0x1F
			fallthrough
		case 2:
			buf[3] |= (src[1] << 4) & 0x1F
			buf[2] = (src[1] >> 1) & 0x1F
			buf[1] = (src[1] >> 6) & 0x1F
			fallthrough
		case 1:
			buf[1] |= (src[0] << 2) & 0x1F
			buf[0] = src[0] >> 3
		}

		// Encode 5-bit blocks using the base32 alphabet
		size := len(dst)
		if size >= 8 {
			// Common case, unrolled for extra performance
			dst[0] = b.encodeMap[buf[0]&31]
			dst[1] = b.encodeMap[buf[1]&31]
			dst[2] = b.encodeMap[buf[2]&31]
			dst[3] = b.encodeMap[buf[3]&31]
			dst[4] = b.encodeMap[buf[4]&31]
			dst[5] = b.encodeMap[buf[5]&31]
			dst[6] = b.encodeMap[buf[6]&31]
			dst[7] = b.encodeMap[buf[7]&31]
		} else {
			for i := 0; i < size; i++ {
				dst[i] = b.encodeMap[buf[i]&31]
			}
		}

		// Pad the final quantum
		if len(src) < 5 {
			if b.padding == base2.NotPadding {
				break
			}

			dst[7] = byte(b.padding)
			if len(src) < 4 {
				dst[6] = byte(b.padding)
				dst[5] = byte(b.padding)
				if len(src) < 3 {
					dst[4] = byte(b.padding)
					if len(src) < 2 {
						dst[3] = byte(b.padding)
						dst[2] = byte(b.padding)
					}
				}
			}

			break
		}

		src = src[5:]
		dst = dst[8:]
	}
}

func (b *base32Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, b.encodedLen(len(src)))
	b.encode(dst, src)
	return dst, nil
}

func (b *base32Codec) decodedLen(n int) int {
	if b.padding == base2.NotPadding {
		return n * 5 / 8
	}
	return n / 8 * 5
}

func (b *base32Codec) decode(dst, src []byte) (n int, end bool, err error) {
	dsti := 0
	olen := len(src)

	for len(src) > 0 && !end {
		// Decode quantum using the base32 alphabet
		var dbuf [8]byte
		dlen := 8

		for j := 0; j < 8; {

			if len(src) == 0 {
				if b.padding != base2.NotPadding {
					// We have reached the end and are missing padding
					pos := olen - len(src) - j
					return n, false, base2.ErrEncodedText(codec, src[pos], pos)
				}
				// We have reached the end and are not expecting any padding
				dlen, end = j, true
				break
			}
			in := src[0]
			src = src[1:]
			if in == byte(b.padding) && j >= 2 && len(src) < 8 {
				// We've reached the end and there's padding
				if len(src)+j < 8-1 {
					// not enough padding
					return n, false, base2.ErrEncodedText(codec, src[olen], olen)
				}
				for k := 0; k < 8-1-j; k++ {
					if len(src) > k && src[k] != byte(b.padding) {
						// incorrect padding
						pos := olen - len(src) + k - 1
						return n, false, base2.ErrEncodedText(codec, src[pos], pos)
					}
				}
				dlen, end = j, true
				// 7, 5 and 2 are not valid padding lengths, and so 1, 3 and 6 are not
				// valid dlen values. See RFC 4648 Section 6 "base 32 Encoding" listing
				// the five valid padding lengths, and Section 9 "Illustrations and
				// Examples" for an illustration for how the 1st, 3rd and 6th base32
				// src bytes do not yield enough information to decode a dst byte.
				if dlen == 1 || dlen == 3 || dlen == 6 {
					pos := olen - len(src) - 1
					return n, false, base2.ErrEncodedText(codec, src[pos], pos)
				}
				break
			}
			elem, ok := b.decodeMap[in]
			if !ok {
				pos := olen - len(src) - 1
				return n, false, base2.ErrEncodedText(codec, src[pos], pos)
			}
			dbuf[j] = byte(elem)

			j++
		}

		// Pack 8x 5-bit source blocks into 5 byte destination
		// quantum
		switch dlen {
		case 8:
			dst[dsti+4] = dbuf[6]<<5 | dbuf[7]
			n++
			fallthrough
		case 7:
			dst[dsti+3] = dbuf[4]<<7 | dbuf[5]<<2 | dbuf[6]>>3
			n++
			fallthrough
		case 5:
			dst[dsti+2] = dbuf[3]<<4 | dbuf[4]>>1
			n++
			fallthrough
		case 4:
			dst[dsti+1] = dbuf[1]<<6 | dbuf[2]<<1 | dbuf[3]>>4
			n++
			fallthrough
		case 2:
			dst[dsti+0] = dbuf[0]<<3 | dbuf[1]>>2
			n++
		}
		dsti += 5
	}
	return n, end, nil
}

func (b *base32Codec) Decode(src []byte) ([]byte, error) {
	if len(src) <= 0 {
		return []byte{}, nil
	}
	dst := make([]byte, b.decodedLen(len(src)))
	n, _, err := b.decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
