// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/18 22:41:30                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base62                                                                                                      *
// * File: base62_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base62

import (
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
			name:      "base62-1",
			plainText: []byte("this is encode."),
			encodText: []byte("rCJBhfvjUf12Ocy2NNtO"),
		},
		{
			name:      "base62-2",
			plainText: []byte("this is base62 encode."),
			encodText: []byte("4ZAp5fLvpSOmQjtik54pZatNTlTiik"),
		},
		{
			name:      "base62-3",
			plainText: []byte("这是一次 base62 编码/解码测试。"),
			encodText: []byte("5NZ7edvjXxa5sRWzZTZrJDqdcQo7zNxNLtzptPTTpMWAjhYgtER78b7kA"),
		},
	}
	encoder     = "@#$%&*(){}ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	cusCodec, _ = NewCodec(encoder)
	testsCus    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base62-1",
			plainText: []byte("this is encode."),
			encodText: []byte("rCJBhfvjUf#$Ocy$NNtO"),
		},
		{
			name:      "base62-2",
			plainText: []byte("this is base62 encode."),
			encodText: []byte("&ZAp*fLvpSOmQjtik*&pZatNTlTiik"),
		},
		{
			name:      "base62-3",
			plainText: []byte("这是一次 base62 编码/解码测试。"),
			encodText: []byte("*NZ)edvjXxa*sRWzZTZrJDqdcQo)zNxNLtzptPTTpMWAjhYgtER){b)kA"),
		},
	}
)

func TestBase62StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base62.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base62.Encode() failed!")
			} else {
				t.Log("base62.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase62StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base62.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base62.Decode() failed!")
			} else {
				t.Log("base62.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase62CusCodec_Encode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base62.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base62.Encode() failed!")
			} else {
				t.Log("base62.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase62CusCodec_Decode(t *testing.T) {
	for _, tt := range testsCus {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cusCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base62.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base62.Decode() failed!")
			} else {
				t.Log("base62.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase62Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base62
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase62Codec_Encode-16           4216830               289.2 ns/op            24 B/op          1 allocs/op
// BenchmarkBase62Codec_Encode-16           3998264               287.7 ns/op            24 B/op          1 allocs/op
// BenchmarkBase62Codec_Encode-16           4219528               284.4 ns/op            24 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base62       4.630s
func BenchmarkBase62Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(testsStd[0].plainText)
	}
}

// go test -bench=BenchmarkBase62Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base62
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase62Codec_Decode-16           3044401               348.6 ns/op            16 B/op          1 allocs/op
// BenchmarkBase62Codec_Decode-16           2978906               378.0 ns/op            16 B/op          1 allocs/op
// BenchmarkBase62Codec_Decode-16           3385266               337.8 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base62       5.081s
func BenchmarkBase62Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(testsStd[0].encodText)
	}
}
