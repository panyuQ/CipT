// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/25 19:18:37                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base                                                                                                        *
// * File: error.go                                                                                                    *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base

import "github.com/pkg/errors"

func ErrEncoderSize(codec string, size int) error {
	return errors.Errorf("codec/%s: encoding alphabet is not %d-bytes long.", codec, size)
}

func ErrEncoderRepeatChar(codec string) error {
	return errors.Errorf("codec/%s: encoding alphabet has repeat character.", codec)
}

func ErrEncoderIllegalChar(codec string) error {
	return errors.Errorf("codec/%s: encoding alphabet contains illegal character.", codec)
}

func ErrEncodedText(codec string, char byte, pos int) error {
	return errors.Errorf("codec/%s: decode src include not in encoder character: %+v, pos: %d.", codec, char, pos)
}

func ErrEncodedTextSize(codec string, size, block int) error {
	return errors.Errorf("codec/%s: invalid decode src size:%d, it should be multiple of %d.", codec, size, block)
}

func ErrEncodedTextMod(codec, mods string, size, mod int) error {
	return errors.Errorf("codec/%s: invalid decode data length: %d, It should be mod = %s not %d.", codec, size, mods, mod)
}

func ErrPaddingIllegalChar(codec string) error {
	return errors.Errorf("codec/%s: padding contained illegal character.", codec)
}

func ErrPaddingAlphabet(codec string) error {
	return errors.Errorf("codec/%s: padding contained in alphabet.", codec)
}

func ErrDecodeSrcDataSize(codec string, size int) error {
	return errors.Errorf("codec/%s: invalid decode src data size: %d.", codec, size)
}
