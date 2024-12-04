// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/29 22:39:12                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base92                                                                                                      *
// * File: base92_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base92

import (
	"fmt"
	"testing"
)

func TestBase92Codec_Encode(t *testing.T) {
	src := []byte{0, 0, 0, 69, 68, 64, 79, 99}
	std, err := NewCodec(stdEncoder)
	if err != nil {
		t.Error(err)
		return
	}

	dst, err := std.Encode(src)
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
	fmt.Println(dst2)

}

//8AArJ45Yvcrsx7XWiJkC2JiSfZAPo
//d.PHS$KQPK}R)cHr*(xebO&V7q
//d.PHS$KQPK}R)cHr*(xebO&V7q
