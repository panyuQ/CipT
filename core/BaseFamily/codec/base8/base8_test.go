// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/22 23:47:56                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base8                                                                                                       *
// * File: base8_test.go                                                                                               *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base8

import (
	"CipT/core/BaseFamily/codec/base"
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
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("3506415134620151346201453346155731062456"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("35064151346201513462014230271545160201453346155731062456"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("72137631715142577113420071526241100611413466247010163674455636404022775051721747501007465530575053712743401010=="),
		},
	}
	codecr, _ = NewCodec("ABCDEFGH", base.StdPadding)
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base8-1",
			plainText: []byte("this is encode."),
			encodText: []byte("DFAGEBFBDEGCABFBDEGCABEFDDEGBFFHDBAGCEFG"),
		},
		{
			name:      "base8-2",
			plainText: []byte("this is base8 encode."),
			encodText: []byte("DFAGEBFBDEGCABFBDEGCABECDACHBFEFBGACABEFDDEGBFFHDBAGCEFG"),
		},
		{
			name:      "base8-3",
			plainText: []byte("这是一次 base8 编码/解码测试。"),
			encodText: []byte("HCBDHGDBHBFBECFHHBBDECAAHBFCGCEBBAAGBBEBDEGGCEHABABGDGHEEFFGDGEAEACCHHFAFBHCBHEHFABAAHEGFFDAFHFAFDHBCHEDEABABA=="),
		},
	}
)

func TestBase8Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base8.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base8.Encode() failed!")
			} else {
				t.Log("base8.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase8Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base8.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base8.Decode() failed!")
			} else {
				t.Log("base8.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase8CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base8.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base8.Encode() failed!")
			} else {
				t.Log("base8.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase8CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base8.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base8.Decode() failed!")
			} else {
				t.Log("base8.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase8Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base8
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase8Codec_Encode-16           30598740                37.39 ns/op           48 B/op          1 allocs/op
// BenchmarkBase8Codec_Encode-16           31586095                37.80 ns/op           48 B/op          1 allocs/op
// BenchmarkBase8Codec_Encode-16           32057018                38.28 ns/op           48 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base8        3.999s
func BenchmarkBase8Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

// go test -bench=BenchmarkBase8Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base8
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase8Codec_Decode-16            3145015               378.9 ns/op            16 B/op          1 allocs/op
// BenchmarkBase8Codec_Decode-16            3178938               372.4 ns/op            16 B/op          1 allocs/op
// BenchmarkBase8Codec_Decode-16            3179215               377.5 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base8        4.962s
func BenchmarkBase8Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
