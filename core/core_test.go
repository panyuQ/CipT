package core

import (
	"fmt"
	"testing"
)

func TestGetMethods(t *testing.T) {
	fmt.Println(GetMethods(true))
	fmt.Println(GetMethods(false))
}

func TestXXEncode_Encode(t *testing.T) {
	text := "你好2311111111111111111111111111111231111111111111111111111111111"
	want := map[string]string{
		"UTF-8":   "ht9qUtOKxAXAlAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH6nAH2lAH2l\nKAH2lAH2lAH2lAH2lAH2lAH2lAH2lAE++",
		"GBK":     "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
		"GB2312":  "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
		"GB18030": "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
	}
	for charset, expected := range want {
		cipT := CipT{
			Encoding: charset,
			Method:   "XXEncode",
		}
		result, err := cipT.Encode([]string{text})
		if err != nil {
			t.Errorf("(%s) Failed to encode: %v", charset, err)
			continue
		}

		if result[0] != expected {
			t.Errorf("(%s) Mismatched result: Got: %s, Want: %s", charset, result, expected)
		} else {
			t.Logf("(%s) Successfully encoded (Got: %s)", charset, result)
		}
	}
}
func TestXXEncode_Decode(t *testing.T) {
	text := "你好2311111111111111111111111111111231111111111111111111111111111"
	want := map[string]string{
		"UTF-8":   "ht9qUtOKxAXAlAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH6nAH2lAH2l\nKAH2lAH2lAH2lAH2lAH2lAH2lAH2lAE++",
		"GBK":     "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
		"GB2312":  "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
		"GB18030": "hlCCukn6nAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2lAH2mAn2lAH2lAH2l\nIAH2lAH2lAH2lAH2lAH2lAH2lAH2+",
	}
	for charset, expected := range want {
		cipT := CipT{
			Encoding: charset,
			Method:   "XXEncode",
		}
		result, err := cipT.Decode([]string{expected})
		if err != nil {
			t.Errorf("(%s) Failed to encode: %v", charset, err)
			continue
		}

		if result[0] != text {
			t.Errorf("(%s) Mismatched result: Got: %s, Want: %s", charset, result, expected)
		} else {
			t.Logf("(%s) Successfully encoded (Got: %s)", charset, result)
		}
	}
}

func TestUUEncode_Encode(t *testing.T) {
	text := "你好2311111111111111111111111111111231111111111111111111111111111"
	want := map[string]string{
		"UTF-8":   "MY+V@Y:6],C,Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3(S,3$Q,3$Q\n6,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,0``",
		"GBK":     "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
		"GB2312":  "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
		"GB18030": "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
	}
	for charset, expected := range want {
		cipT := CipT{
			Encoding: charset,
			Method:   "UUEncode",
		}
		result, err := cipT.Encode([]string{text})
		if err != nil {
			t.Errorf("(%s) Failed to encode: %v", charset, err)
			continue
		}

		if result[0] != expected {
			t.Errorf("(%s) Mismatched result: Got: %s, Want: %s", charset, result, expected)
		} else {
			t.Logf("(%s) Successfully encoded (Got: %s)", charset, result)
		}
	}
}
func TestUUEncode_Decode(t *testing.T) {
	text := "你好2311111111111111111111111111111231111111111111111111111111111"
	want := map[string]string{
		"UTF-8":   "MY+V@Y:6],C,Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3(S,3$Q,3$Q\n6,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,0``",
		"GBK":     "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
		"GB2312":  "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
		"GB18030": "MQ..ZPS(S,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$R,S$Q,3$Q,3$Q\n4,3$Q,3$Q,3$Q,3$Q,3$Q,3$Q,3$`",
	}
	for charset, expected := range want {
		cipT := CipT{
			Encoding: charset,
			Method:   "UUEncode",
		}
		result, err := cipT.Decode([]string{expected})
		if err != nil {
			t.Errorf("(%s) Failed to encode: %v", charset, err)
			continue
		}

		if result[0] != text {
			t.Errorf("(%s) Mismatched result: Got: %s, Want: %s", charset, result, expected)
		} else {
			t.Logf("(%s) Successfully encoded (Got: %s)", charset, result)
		}
	}
}
