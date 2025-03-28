// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 21:53:50                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base100                                                                                                     *
// * File: base100.go                                                                                                  *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base100

import (
	"CipT/core/BaseFamily/codec/internal/errs"
)

const (
	codec = "base100"
)

type base100Codec struct{}

var StdCodec base100Codec

//func NewCodec() *base100Codec {
//	return &base100Codec{}
//}

func (b *base100Codec) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, len(src)*4)
	for k, v := range src {
		dst[k*4+0] = 0xF0
		dst[k*4+1] = 0x9F
		dst[k*4+2] = byte((uint16(v)+55)/64 + 0x8F)
		dst[k*4+3] = (v+55)%64 + 0x80
	}
	return dst, nil
}

func (b *base100Codec) Decode(src []byte) ([]byte, error) {
	if len(src)%4 != 0 {
		return nil, errs.ErrEncodedTextSize(codec, len(src), 4)
	}
	dst := make([]byte, len(src)/4)
	for k := 0; k != len(src); k += 4 {
		if src[k] != 0xF0 {
			return nil, errs.ErrEncodedText(codec, src[k], k)
		}
		if src[k+1] != 0x9F {
			return nil, errs.ErrEncodedText(codec, src[k+1], k+1)
		}
		dst[k/4] = (src[k+2]-0x8F)*64 + src[k+3] - 0x80 - 55
	}
	return dst, nil
}
