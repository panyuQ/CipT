// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/10 22:36:52                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base16                                                                                                      *
// * File: base16_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base16

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
			name:      "base16-1",
			plainText: []byte("this is encode."),
			encodText: []byte("7468697320697320656e636f64652e"),
		},
		{
			name:      "base16-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("7468697320697320626173653820656e636f64652e"),
		},
		{
			name:      "base16-3",
			plainText: []byte("这是一次 base16 编码/解码测试。"),
			encodText: []byte("e8bf99e698afe4b880e6aca12062617365313620e7bc96e7a0812fe8a7a3e7a081e6b58be8af95e38082"),
		},
	}
	codecr, _ = NewCodec("ABCDEFGHIJKLMNOP")
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("HEGIGJHDCAGJHDCAGFGOGDGPGEGFCO"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("HEGIGJHDCAGJHDCAGCGBHDGFDICAGFGOGDGPGEGFCO"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("OILPJJOGJIKPOELIIAOGKMKBCAGCGBHDGFDICAOHLMJGOHKAIBCPOIKHKDOHKAIBOGLFILOIKPJFODIAIC"),
		},
	}
)

func TestBase16Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base16.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base16.Encode() failed!")
			} else {
				t.Log("base16.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase16Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base16.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base16.Decode() failed!")
			} else {
				t.Log("base16.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase16CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base16.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base16.Encode() failed!")
			} else {
				t.Log("base16.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase16CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base16.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base16.Decode() failed!")
			} else {
				t.Log("base16.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase16Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base16
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase16Codec_Encode-16          38236652                26.84 ns/op           32 B/op          1 allocs/op
// BenchmarkBase16Codec_Encode-16          43396341                26.12 ns/op           32 B/op          1 allocs/op
// BenchmarkBase16Codec_Encode-16          45122450                26.20 ns/op           32 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base16       3.737s
func BenchmarkBase16Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

// go test -bench=BenchmarkBase16Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base16
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase16Codec_Decode-16           4313695               259.4 ns/op            16 B/op          1 allocs/op
// BenchmarkBase16Codec_Decode-16           4624881               259.1 ns/op            16 B/op          1 allocs/op
// BenchmarkBase16Codec_Decode-16           4647886               257.5 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base16       4.620s
func BenchmarkBase16Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
