// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 23:34:00                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base91                                                                                                      *
// * File: base91_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base91

import (
	"fmt"
	"testing"
)

func TestBase91Codec_Encode(t *testing.T) {
	src := []byte("this is base91 encoded.")
	dst, err := StdCodec.Encode(src)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(dst))
	dst2, err := StdCodec.Decode(dst)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(dst2))
}
