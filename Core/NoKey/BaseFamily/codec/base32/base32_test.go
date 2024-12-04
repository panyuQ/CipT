// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/23 23:45:19                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base32                                                                                                      *
// * File: base32_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base32

import (
	"CipT/Core/NoKey/BaseFamily/codec/base"
	"reflect"
	"testing"
)

var (
	testsStd = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("ORUGS4ZANFZSAZLOMNXWIZJO"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("ORUGS4ZANFZSAYTBONSTGMRAMVXGG33EMUXA===="),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("5C7ZTZUYV7SLRAHGVSQSAYTBONSTGMRA466JNZ5AQEX6RJ5D46QIDZVVRPUK7FPDQCBA===="),
		},
	}
	testsHex = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("EHK6ISP0D5PI0PBECDNM8P9E"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("EHK6ISP0D5PI0OJ1EDIJ6CH0CLN66RR4CKN0===="),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("T2VPJPKOLVIBH076LIGI0OJ1EDIJ6CH0SUU9DPT0G4NUH9T3SUG83PLLHFKAV5F3G210===="),
		},
	}
	testsRawStd = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("ORUGS4ZANFZSAZLOMNXWIZJO"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("ORUGS4ZANFZSAYTBONSTGMRAMVXGG33EMUXA"),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("5C7ZTZUYV7SLRAHGVSQSAYTBONSTGMRA466JNZ5AQEX6RJ5D46QIDZVVRPUK7FPDQCBA"),
		},
	}
	testsRawHex = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("EHK6ISP0D5PI0PBECDNM8P9E"),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("EHK6ISP0D5PI0OJ1EDIJ6CH0CLN66RR4CKN0"),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("T2VPJPKOLVIBH076LIGI0OJ1EDIJ6CH0SUU9DPT0G4NUH9T3SUG83PLLHFKAV5F3G210"),
		},
	}
	encoder     = "!@#$%^&*()_+:{[}]<>?/-~123456789"
	cusCodec, _ = NewCodec(encoder, base.StdPadding)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base32-1",
			plainText: []byte("this is encode."),
			encodText: []byte("[</&>63!{^3>!3+[:{1~(3)["),
		},
		{
			name:      "base32-2",
			plainText: []byte("this is base32 encode."),
			encodText: []byte("[</&>63!{^3>!2?@[{>?&:<!:-1&&55%:/1!===="),
		},
		{
			name:      "base32-3",
			plainText: []byte("这是一次 base32 编码/解码测试。"),
			encodText: []byte("7#93?3/2-9>+<!*&->]>!2?@[{>?&:<!688){37!]%18<)7$68]($3--<}/_9^}$]#@!===="),
		},
	}
)

func TestBase32StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase32HexCodec_Encode(t *testing.T) {
	for _, tt := range testsHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32HexCodec_Decode(t *testing.T) {
	for _, tt := range testsHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase32RawStdCodec_Encode(t *testing.T) {
	for _, tt := range testsRawStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawStdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32RawStdCodec_Decode(t *testing.T) {
	for _, tt := range testsRawStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawStdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase32RawHexCodec_Encode(t *testing.T) {
	for _, tt := range testsRawHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawHexCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32RawHexCodec_Decode(t *testing.T) {
	for _, tt := range testsRawHex {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RawHexCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase32CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base32.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base32.Encode() failed!")
			} else {
				t.Log("base32.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase32CusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base32.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base32.Decode() failed!")
			} else {
				t.Log("base32.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase32Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base32
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase32Codec_Encode-16          36709371                29.38 ns/op           24 B/op          1 allocs/op
// BenchmarkBase32Codec_Encode-16          37686193                29.37 ns/op           24 B/op          1 allocs/op
// BenchmarkBase32Codec_Encode-16          41128285                30.09 ns/op           24 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base32       3.808s
func BenchmarkBase32Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(testsStd[0].plainText)
	}
}

// go test -bench=BenchmarkBase32Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base32
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase32Codec_Decode-16           3631639               336.2 ns/op            16 B/op          1 allocs/op
// BenchmarkBase32Codec_Decode-16           3395719               335.7 ns/op            16 B/op          1 allocs/op
// BenchmarkBase32Codec_Decode-16           3612626               331.6 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base32       4.891s
func BenchmarkBase32Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(testsStd[0].encodText)
	}
}
