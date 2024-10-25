// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 21:34:10                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: ascii85                                                                                                     *
// * File: ascii85_test.go                                                                                             *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package ascii85

import (
	"reflect"
	"testing"
)

var (
	tests = []struct {
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

func TestAscii85Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("ascii85.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("ascii85.Encode() failed!")
			} else {
				t.Log("ascii85.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestAscii85Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("ascii85.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Log("ascii85.Decode() failed!")
			} else {
				t.Log("ascii85.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkAscii85Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/ascii85
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkAscii85Codec_Encode-16         30266190                38.81 ns/op           24 B/op          1 allocs/op
// BenchmarkAscii85Codec_Encode-16         30370827                38.72 ns/op           24 B/op          1 allocs/op
// BenchmarkAscii85Codec_Encode-16         28559324                38.80 ns/op           24 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/ascii85      5.494s
func BenchmarkAscii85Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

// go test -bench=BenchmarkAscii85Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/ascii85
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkAscii85Codec_Decode-16         34787750                35.07 ns/op           24 B/op          1 allocs/op
// BenchmarkAscii85Codec_Decode-16         34775451                35.25 ns/op           24 B/op          1 allocs/op
// BenchmarkAscii85Codec_Decode-16         33139100                34.97 ns/op           24 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/ascii85      5.735s
func BenchmarkAscii85Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
