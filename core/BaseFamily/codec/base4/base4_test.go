// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 22:34:18                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base4                                                                                                       *
// * File: base4_test.go                                                                                               *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base4

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
			name:      "base4-1",
			plainText: []byte("this is encode."),
			encodText: []byte("131012201221130302001221130302001211123212031233121012110232"),
		},
		{
			name:      "base4-2",
			plainText: []byte("this is base4 encode."),
			encodText: []byte("131012201221130302001221130302001202120113031211031002001211123212031233121012110232"),
		},
		{
			name:      "base4-3",
			plainText: []byte("这是一次 base4 编码/解码测试。"),
			encodText: []byte("3220233321213212212022333210232020003212223022010200120212011303121103100200321323302" +
				"1123213220020010233322022132203321322002001321223112023322022332111320320002002"),
		},
	}
	codecr, _ = NewCodec("ABCD")
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base4-1",
			plainText: []byte("this is encode."),
			encodText: []byte("BDBABCCABCCBBDADACAABCCBBDADACAABCBBBCDCBCADBCDDBCBABCBBACDC"),
		},
		{
			name:      "base4-2",
			plainText: []byte("this is base4 encode."),
			encodText: []byte("BDBABCCABCCBBDADACAABCCBBDADACAABCACBCABBDADBCBBADBAACAABCBBBCDCBCADBCDDBCBABCBBACDC"),
		},
		{
			name:      "base4-3",
			plainText: []byte("这是一次 base4 编码/解码测试。"),
			encodText: []byte("DCCACDDDCBCBDCBCCBCACCDDDCBACDCACAAADCBCCCDACCABACAABCACBCABBDADBCBBADBAACAADCBDCD" +
				"DACBBCDCBDCCAACAABACDDDCCACCBDCCADDCBDCCAACAABDCBCCDBBCACDDCCACCDDCBBBDCADCAAACAAC"),
		},
	}
)

func TestBase4Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base4.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base4.Encode() failed!")
			} else {
				t.Log("base4.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase4Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base4.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base4.Decode() failed!")
			} else {
				t.Log("base4.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase4CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base4.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base4.Encode() failed!")
			} else {
				t.Log("base4.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase4CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base4.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base4.Decode() failed!")
			} else {
				t.Log("base4.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase4Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base4
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase4Codec_Encode-16           28809379                41.91 ns/op           64 B/op          1 allocs/op
// BenchmarkBase4Codec_Encode-16           28452806                42.38 ns/op           64 B/op          1 allocs/op
// BenchmarkBase4Codec_Encode-16           27547576                42.74 ns/op           64 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base4        3.970s
func BenchmarkBase4Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

// go test -bench=BenchmarkBase4Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base4
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase4Codec_Decode-16            1449396               817.4 ns/op            16 B/op          1 allocs/op
// BenchmarkBase4Codec_Decode-16            1465980               814.9 ns/op            16 B/op          1 allocs/op
// BenchmarkBase4Codec_Decode-16            1432980               812.8 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base4        6.283s
func BenchmarkBase4Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
