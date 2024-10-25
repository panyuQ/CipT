// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/12/14 22:21:12                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base24                                                                                                      *
// * File: base64_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base24

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
			name:      "base24-1",
			plainText: []byte("this is encode."),
			encodText: []byte("H3ETZ75C9FZAG44W82WYF4STR7CP"),
		},
		{
			name:      "base24-2",
			plainText: []byte("this is base24 encode."),
			encodText: []byte("H3ETZ75C9FZAG447EWP6WBGRP6GKGP6C78T4WBHPZZ"),
		},
		{
			name:      "base24-3",
			plainText: []byte("这是一次 base24 编码/解码测试。"),
			encodText: []byte("9HG6G86KG8AK7P5F6E29AC936KC24W3E6EP94EH5XY6CBGBGP6833BEP5GPSXWS9HEH5T25ER4W4Z"),
		},
	}
	codecr, _ = NewCodec("0123456789ABCDEFGHIJKLMN")
	tests1    = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base24-1",
			plainText: []byte("this is encode."),
			encodText: []byte("A56C0FB2K701988LH3LN78JCIF2G"),
		},
		{
			name:      "base24-2",
			plainText: []byte("this is base24 encode."),
			encodText: []byte("A56C0FB2K701988F6LGEL49IGE9D9GE2FHC8L4AG00"),
		},
		{
			name:      "base24-3",
			plainText: []byte("这是一次 base24 编码/解码测试。"),
			encodText: []byte("KA9E9HED9H1DFGB7E63K12K5ED238L56E6GK86ABMNE24949GEH5546GB9GJMLJKA6ABC3B6I8L80"),
		},
	}
)

func TestBase24Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base24.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base24.Encode() failed!")
			} else {
				t.Log("base24.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase24Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base24.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base24.Decode() failed!")
			} else {
				t.Log("base24.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase24CusCodec_Encode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base24.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base24.Encode() failed!")
			} else {
				t.Log("base24.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase24CusCodec_Decode(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecr.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base24.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base24.Decode() failed!")
			} else {
				t.Log("base24.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

// go test -bench=BenchmarkBase24Codec_Encode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base24
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase24Codec_Encode-16          17161538                66.69 ns/op           48 B/op          2 allocs/op
// BenchmarkBase24Codec_Encode-16          19548144                61.38 ns/op           48 B/op          2 allocs/op
// BenchmarkBase24Codec_Encode-16          19138449                61.29 ns/op           48 B/op          2 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base24       3.981s
func BenchmarkBase24Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

// go test -bench=BenchmarkBase24Codec_Decode -benchmem -count=3
// goos: windows
// goarch: amd64
// pkg: github.com/caijunjun/codec/base24
// cpu: 12th Gen Intel(R) Core(TM) i7-12650H
// BenchmarkBase24Codec_Decode-16           4828347               244.9 ns/op            16 B/op          1 allocs/op
// BenchmarkBase24Codec_Decode-16           4914670               246.9 ns/op            16 B/op          1 allocs/op
// BenchmarkBase24Codec_Decode-16           4638578               247.0 ns/op            16 B/op          1 allocs/op
// PASS
// ok      github.com/caijunjun/codec/base24       4.597s
func BenchmarkBase24Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
