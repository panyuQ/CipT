package NoKey

import (
	"CipT/NoKey/Base-variant"
	"CipT/codec/ascii85"
	"CipT/codec/base100"
	"CipT/codec/base16"
	"CipT/codec/base24"
	"CipT/codec/base32"
	"CipT/codec/base36"
	"CipT/codec/base4"
	"CipT/codec/base45"
	"CipT/codec/base58"
	"CipT/codec/base64"
	"CipT/codec/base8"
	"CipT/codec/base85"
	"CipT/codec/base91"
	"CipT/codec/base92"
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/url"
)

type Cip struct {
	Before []byte // Before 编码前/解码后
	After  []byte // After 编码后/解码前
	Method string // Method 方法
}

var allMethods = []string{
	"ASCII85",

	"Base4", "Base8", "Base16", "Base32", "Base64",
	"Base24", "Base36", "Base45", "Base58", "Base85", "Base91", "Base92", "Base100",

	"UUEncode", "XXEncode",

	"URL",
}

// enFunc 定义编码方法类型（[]byte）
type enFunc func(data []byte) ([]byte, error)

// deFunc 定义解码方法类型（[]byte）
type deFunc func(data []byte) ([]byte, error)

// 创建一个映射，将 method 值与相应的编码和解码函数关联
var encoderFunc = map[string]enFunc{

	"ASCII85": ascii85.StdCodec.Encode,

	//"Base2":   base2.StdCodec.Encode,
	"Base4":   base4.StdCodec.Encode,
	"Base8":   base8.StdCodec.Encode,
	"Base16":  base16.StdCodec.Encode,
	"Base24":  base24.StdCodec.Encode,
	"Base32":  base32.StdCodec.Encode,
	"Base36":  base36.StdCodec.Encode,
	"Base45":  base45.StdCodec.Encode,
	"Base58":  base58.StdCodec.Encode,
	"Base64":  base64.StdCodec.Encode,
	"Base85":  base85.StdCodec.Encode,
	"Base91":  base91.StdCodec.Encode,
	"Base92":  base92.StdCodec.Encode,
	"Base100": base100.StdCodec.Encode,

	"UUEncode": Base_variant.UUEncode,
	"XXEncode": Base_variant.XXEncode,

	"URL": func(data []byte) ([]byte, error) {
		return []byte(url.QueryEscape(string(data))), nil
	},
}

var decoderFunc = map[string]deFunc{

	"ASCII85": ascii85.StdCodec.Decode,

	//"Base2":   base2.StdCodec.Decode,
	"Base4":   base4.StdCodec.Decode,
	"Base8":   base8.StdCodec.Decode,
	"Base16":  base16.StdCodec.Decode,
	"Base24":  base24.StdCodec.Decode,
	"Base32":  base32.StdCodec.Decode,
	"Base36":  base36.StdCodec.Decode,
	"Base45":  base45.StdCodec.Decode,
	"Base58":  base58.StdCodec.Decode,
	"Base64":  base64.StdCodec.Decode,
	"Base85":  base85.StdCodec.Decode,
	"Base91":  base91.StdCodec.Decode,
	"Base92":  base92.StdCodec.Decode,
	"Base100": base100.StdCodec.Decode,

	"UUEncode": Base_variant.UUDecode,
	"XXEncode": Base_variant.XXDecode,

	"URL": func(data []byte) ([]byte, error) {
		res, err := url.QueryUnescape(string(data))
		return []byte(res), err
	},
}

// FromUTF8_To 把 UTF-8 的数据，转换为 指定编码的数据
var FromUTF8_To = map[string]enFunc{
	"GBK": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewEncoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"GB2312": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewEncoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"GB18030": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewEncoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"HZGB2312": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.HZGB2312.NewEncoder())
		res, err := io.ReadAll(r)
		return res, err
	},
}

// ToUTF8_From 把指定编码的数据，转换为 UTF-8 的数据
var ToUTF8_From = map[string]deFunc{
	"GBK": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"GB2312": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"GB18030": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())
		res, err := io.ReadAll(r)
		return res, err
	},
	"HZGB2312": func(data []byte) ([]byte, error) {
		r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.HZGB2312.NewDecoder())
		res, err := io.ReadAll(r)
		return res, err
	},
}

// GetSupportedMethods 查看支持的加密解密方法
func GetSupportedMethods() []string {
	return allMethods
}

func (t *Cip) Encode() error {
	if encodeFunc, ok := encoderFunc[t.Method]; ok {
		// 如果 method 对应的编码函数存在，则调用
		res, err := encodeFunc(t.Before)
		if err != nil {
			return err
		}
		t.After = res
		return nil
	} else {
		return errors.New("Unknown encoding method ( " + t.Method + " )")
	}
}

func (t *Cip) Decode() error {
	if decodeFunc, ok := decoderFunc[t.Method]; ok {
		// 如果 method 对应的解码函数存在，则调用
		res, err := decodeFunc(t.After)
		if err != nil {
			return err
		}
		t.Before = res
		return nil
	} else {
		return errors.New("Unknown decoding method ( " + t.Method + " )")
	}
}
