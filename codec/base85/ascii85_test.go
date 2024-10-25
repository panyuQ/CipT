// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 22:31:19                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: ascii85                                                                                                     *
// * File: ascii85_test.go                                                                                             *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base85

import (
	"reflect"
	"testing"
)

var (
	encoder, _ = NewCodec(StdEncoder)
	tests      = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "ascii85-1",
			plainText: []byte("this is encode."),
			encodText: []byte("FD,B0+DGm>ASu!rA7[@"),
		},
		{
			name:      "ascii85-2",
			plainText: []byte("this is ascii85 encode."),
			encodText: []byte("FD,B0+DGm>@<5pmBfIsmASu!rA7[@"),
		},
		{
			name:      "ascii85-3",
			plainText: []byte("这是一次 ascii85 编码/解码测试。"),
			encodText: []byte("keEPJR'5S\\JEEr,+CT>$Bk]Oa+QpD'kFdAZkbk3=TRl75Ms.@0j+)^"),
		},
	}
)

func TestAscii85_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encoder.Encode(tt.plainText)
			if err != nil {
				t.Errorf("ascii85.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("ascii85.Encode() success!")
				t.Errorf(" got: %v", got)
				t.Errorf("want: %v", tt.encodText)
			} else {
				t.Log("ascii85.Encode() success!")
				t.Logf(" got: %v", got)
				t.Logf("want: %v", tt.encodText)
			}
		})
	}
}

func TestAscii85_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encoder.Decode(tt.encodText)
			if err != nil {
				t.Errorf("ascii85.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Log("ascii85.Decode() success!")
				t.Errorf(" got: %v", got)
				t.Errorf("want: %v", tt.plainText)
			} else {
				t.Log("ascii85.Decode() success!")
				t.Logf(" got: %v", got)
				t.Logf("want: %v", tt.plainText)
			}
		})
	}
}

// go test -bench=BenchmarkAscii85_Encode -benchmem -count 3
func BenchmarkAscii85_Encode(b *testing.B) {
	text := []byte("this is ascii85 encode.")
	for i := 0; i < b.N; i++ {
		encoder.Encode(text)
	}
}

// go test -bench=BenchmarkAscii85_Decode -benchmem -count 3
func BenchmarkAscii85_Decode(b *testing.B) {
	text := []byte("FD,B0+DGm>@<5pmBfIsmASu!rA7[@")
	for i := 0; i < b.N; i++ {
		encoder.Decode(text)
	}
}
