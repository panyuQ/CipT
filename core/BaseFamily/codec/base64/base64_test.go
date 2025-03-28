// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2024/01/23 22:24:10                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base64                                                                                                      *
// * File: base64_test.go                                                                                              *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base64

import (
	"CipT/core/BaseFamily/codec/base"
	"fmt"
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
			name:      "base64-1",
			plainText: []byte("this is encode."),
			encodText: []byte("dGhpcyBpcyBlbmNvZGUu"),
		},
		{
			name:      "base64-2",
			plainText: []byte("this is base64 encode."),
			encodText: []byte("dGhpcyBpcyBiYXNlNjQgZW5jb2RlLg=="),
		},
		{
			name:      "base64-3",
			plainText: []byte("这是一次 base64 编码/解码测试。"),
			encodText: []byte("6L+Z5piv5LiA5qyhIGJhc2U2NCDnvJbnoIEv6Kej56CB5rWL6K+V44CC"),
		},
	}
	testsUrl = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base64-1",
			plainText: []byte("this is encode."),
			encodText: []byte("dGhpcyBpcyBlbmNvZGUu"),
		},
		{
			name:      "base64-2",
			plainText: []byte("this is base64 encode."),
			encodText: []byte("dGhpcyBpcyBiYXNlNjQgZW5jb2RlLg=="),
		},
		{
			name:      "base64-3",
			plainText: []byte("这是一次 base64 编码/解码测试。"),
			encodText: []byte("6L-Z5piv5LiA5qyhIGJhc2U2NCDnvJbnoIEv6Kej56CB5rWL6K-V44CC"),
		},
	}
	testsStdRaw = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base64-1",
			plainText: []byte("this is encode."),
			encodText: []byte("dGhpcyBpcyBlbmNvZGUu"),
		},
		{
			name:      "base64-2",
			plainText: []byte("this is base64 encode."),
			encodText: []byte("dGhpcyBpcyBiYXNlNjQgZW5jb2RlLg"),
		},
		{
			name:      "base64-3",
			plainText: []byte("这是一次 base64 编码/解码测试。"),
			encodText: []byte("6L+Z5piv5LiA5qyhIGJhc2U2NCDnvJbnoIEv6Kej56CB5rWL6K+V44CC"),
		},
	}
)

func TestBase64StdCodec_Encode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base64.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base64.Encode() failed!")
			} else {
				t.Log("base64.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase64StdCodec_Decode(t *testing.T) {
	for _, tt := range testsStd {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base64.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base64.Decode() failed!")
			} else {
				t.Log("base64.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase64UrlCodec_Encode(t *testing.T) {
	for _, tt := range testsUrl {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UrlCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base64.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base64.Encode() failed!")
			} else {
				t.Log("base64.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase64UrlCodec_Decode(t *testing.T) {
	for _, tt := range testsUrl {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UrlCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base64.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base64.Decode() failed!")
			} else {
				t.Log("base64.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase64StdRawCodec_Encode(t *testing.T) {
	for _, tt := range testsStdRaw {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdRawCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base64.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Error("base64.Encode() failed!")
			} else {
				t.Log("base64.Encode() success!")
			}
			t.Logf(" got: %v", string(got))
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase64StdRawCodec_Decode(t *testing.T) {
	for _, tt := range testsStdRaw {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdRawCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base64.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Error("base64.Decode() failed!")
			} else {
				t.Log("base64.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestXXEncode_Encode(t *testing.T) {
	XXCodec, _ := NewCodec(stdEncoder, base.StdPadding)
	res, err := XXCodec.Encode([]byte("Hello, World!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	if err != nil {
		t.Errorf("XXEncode() error = %v", err)
	} else {
		fmt.Println(string(res))
		fmt.Println(res)
	}
}
