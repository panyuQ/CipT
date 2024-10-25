// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/03 23:50:51                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base92                                                                                                      *
// * File: base92x.go                                                                                                  *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base92

import (
	"CipT/codec/base"
	"CipT/codec/internal/basex"
)

type base92Codecx struct {
	*basex.Basex
}

func NewBase92Codecx(encoder string) (base.IEncoding, error) {
	b, err := basex.NewBasex(codec, encoder, stdEncoderSize)
	if err != nil {
		return nil, err
	}
	return &base92Codecx{b}, nil
}

func (b *base92Codecx) Encode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.Encode(src)
}

func (b *base92Codecx) Deocde(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	return b.Decode(src)
}
