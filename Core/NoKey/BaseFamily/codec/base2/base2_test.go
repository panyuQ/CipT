// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/21 22:04:39                                                                                         *
// * Proj: codec                                                                                                       *
// * Pack: base2                                                                                                       *
// * File: base2_test.go                                                                                               *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base2

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
			name:      "base2-1",
			plainText: []byte("this is encode."),
			encodText: []byte("01110100011010000110100101110011001000000110100101" +
				"1100110010000001100101011011100110001101101111011001000110010100101110"),
		},
		{
			name:      "base2-2",
			plainText: []byte("this is base2 encode."),
			encodText: []byte("0111010001101000011010010111001100100000011010010111" +
				"001100100000011000100110000101110011011001010011001000100000011001" +
				"01011011100110001101101111011001000110010100101110"),
		},
		{
			name:      "base2-3",
			plainText: []byte("这是一次 base2 编码/解码测试。"),
			encodText: []byte("1110100010111111100110011110011010011000101011111110" +
				"010010111000100000001110011010101100101000010010000001100010011000" +
				"010111001101100101001100100010000011100111101111001001011011100111" +
				"101000001000000100101111111010001010011110100011111001111010000010" +
				"000001111001101011010110001011111010001010111110010101111000111000" +
				"000010000010"),
		},
	}
	codecAB, _ = NewCodec("AB")
	tests2     = []struct {
		name      string
		plainText []byte
		encodText []byte
	}{
		{
			name:      "base2-1",
			plainText: []byte("this is encode."),
			encodText: []byte("ABBBABAAABBABAAAABBABAABABBBAABBAABAAAAAABBABAABABBB" +
				"AABBAABAAAAAABBAABABABBABBBAABBAAABBABBABBBBABBAABAAABBAABABAABABBBA"),
		},
		{
			name:      "base2-2",
			plainText: []byte("this is base2 encode."),
			encodText: []byte("ABBBABAAABBABAAAABBABAABABBBAABBAABAAAAAABBABAABABBBA" +
				"ABBAABAAAAAABBAAABAABBAAAABABBBAABBABBAABABAABBAABAAABAAAAAABBAABAB" +
				"ABBABBBAABBAAABBABBABBBBABBAABAAABBAABABAABABBBA"),
		},
		{
			name:      "base2-3",
			plainText: []byte("这是一次 base2 编码/解码测试。"),
			encodText: []byte("BBBABAAABABBBBBBBAABBAABBBBAABBABAABBAAABABABBBBBBBAAB" +
				"AABABBBAAABAAAAAAABBBAABBABABABBAABABAAAABAABAAAAAABBAAABAABBAAAABAB" +
				"BBAABBABBAABABAABBAABAAABAAAAABBBAABBBBABBBBAABAABABBABBBAABBBBABAAA" +
				"AABAAAAAABAABABBBBBBBABAAABABAABBBBABAAABBBBBAABBBBABAAAAABAAAAAABBB" +
				"BAABBABABBABABBAAABABBBBBABAAABABABBBBBAABABABBBBAAABBBAAAAAAABAAAAABA"),
		},
	}
)

func TestBase2Codec_Encode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base2.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("base2.Encode() failed!")
			} else {
				t.Log("base2.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase2Codec_Decode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdCodec.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base2.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Log("base2.Decode() failed!")
			} else {
				t.Log("base2.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func TestBase2Codec_Encode2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecAB.Encode(tt.plainText)
			if err != nil {
				t.Errorf("base2.Encode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.encodText) {
				t.Log("base2.Encode() failed!")
			} else {
				t.Log("base2.Encode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.encodText)
		})
	}
}

func TestBase2Codec_Decode2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecAB.Decode(tt.encodText)
			if err != nil {
				t.Errorf("base2.Decode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.plainText) {
				t.Log("base2.Decode() failed!")
			} else {
				t.Log("base2.Decode() success!")
			}
			t.Logf(" got: %v", got)
			t.Logf("want: %v", tt.plainText)
		})
	}
}

func BenchmarkBase2Codec_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Encode(tests[0].plainText)
	}
}

func BenchmarkBase2Codec_Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdCodec.Decode(tests[0].encodText)
	}
}
