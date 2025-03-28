// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/24 22:41:25                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: codec                                                                                                       *
// * File: codec.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package codec

import (
	"CipT/core/BaseFamily/codec/ascii85"
	"CipT/core/BaseFamily/codec/base"
	"CipT/core/BaseFamily/codec/base16"
	"CipT/core/BaseFamily/codec/base2"
	"CipT/core/BaseFamily/codec/base24"
	"CipT/core/BaseFamily/codec/base32"
	"CipT/core/BaseFamily/codec/base4"
	"CipT/core/BaseFamily/codec/base8"
)

type Codec struct {
	codec base.IEncoding
	err   error
}

type CodecResult struct {
	err error
	dst []byte
}

func (c Codec) EncodeByString(src string) CodecResult {
	if c.err != nil {
		return CodecResult{err: c.err}
	}
	dst, err := c.codec.Encode(base.StringToBytes(src))
	return CodecResult{dst: dst, err: err}
}

func (c Codec) EncodeByBytes(src []byte) CodecResult {
	if c.err != nil {
		return CodecResult{err: c.err}
	}
	dst, err := c.codec.Encode(src)
	return CodecResult{dst: dst, err: err}
}

func (c Codec) DecodeByString(src string) CodecResult {
	if c.err != nil {
		return CodecResult{err: c.err}
	}
	dst, err := c.codec.Decode(base.StringToBytes(src))
	return CodecResult{dst: dst, err: err}
}

func (c Codec) DecodeByBytes(src []byte) CodecResult {
	if c.err != nil {
		return CodecResult{err: c.err}
	}
	dst, err := c.codec.Decode(src)
	return CodecResult{dst: dst, err: err}
}

func (c CodecResult) ToString() (string, error) {
	if c.err != nil || c.dst == nil {
		return "", nil
	}
	return base.BytesToString(c.dst), nil
}

func (c CodecResult) ToBytes() ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}
	if c.dst == nil {
		return []byte{}, nil
	}
	return c.dst, nil
}

func (c CodecResult) Error() error {
	return c.err
}

func UseStdScii85() Codec {
	return Codec{codec: ascii85.StdCodec}
}

func UseStdBase2() Codec {
	return Codec{codec: base2.StdCodec}
}

func UseCusBase2(encoder string) Codec {
	c := Codec{}
	c.codec, c.err = base2.NewCodec(encoder)
	return c
}

func UseStdBase4() Codec {
	return Codec{codec: base4.StdCodec}
}

func UseCusBase4(encoder string) Codec {
	c := Codec{}
	c.codec, c.err = base4.NewCodec(encoder)
	return c
}

func UseStdBase8() Codec {
	return Codec{codec: base8.StdCodec}
}

func UseCusBase8(encoder string, padding rune) Codec {
	c := Codec{}
	c.codec, c.err = base8.NewCodec(encoder, padding)
	return c
}

func UseStdBase16() Codec {
	return Codec{codec: base16.StdCodec}
}

func UseCusBase16(encoder string) Codec {
	c := Codec{}
	c.codec, c.err = base16.NewCodec(encoder)
	return c
}

func UseStdBase24() Codec {
	return Codec{codec: base24.StdCodec}
}

func UseCusBase24(encoder string) Codec {
	c := Codec{}
	c.codec, c.err = base24.NewCodec(encoder)
	return c
}

func UseStdBase32() Codec {
	return Codec{codec: base32.StdCodec}
}

func UseHexBase32() Codec {
	return Codec{codec: base32.HexCodec}
}

func UseRawStdBase32() Codec {
	return Codec{codec: base32.RawStdCodec}
}

func UseRawHexBase32() Codec {
	return Codec{codec: base32.RawHexCodec}
}

func UseCusBase32(encoder string, padding rune) Codec {
	c := Codec{}
	c.codec, c.err = base32.NewCodec(encoder, padding)
	return c
}
