package NoKey

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding"
	"io"
	"net/url"

	"CipT/Core/NoKey/BaseFamily/codec/ascii85"
	"CipT/Core/NoKey/BaseFamily/codec/base100"
	"CipT/Core/NoKey/BaseFamily/codec/base16"
	"CipT/Core/NoKey/BaseFamily/codec/base24"
	"CipT/Core/NoKey/BaseFamily/codec/base32"
	"CipT/Core/NoKey/BaseFamily/codec/base36"
	"CipT/Core/NoKey/BaseFamily/codec/base4"
	"CipT/Core/NoKey/BaseFamily/codec/base45"
	"CipT/Core/NoKey/BaseFamily/codec/base58"
	"CipT/Core/NoKey/BaseFamily/codec/base64"
	"CipT/Core/NoKey/BaseFamily/codec/base8"
	"CipT/Core/NoKey/BaseFamily/codec/base85"
	"CipT/Core/NoKey/BaseFamily/codec/base91"
	"CipT/Core/NoKey/BaseFamily/codec/base92"
	BaseVariant "CipT/Core/NoKey/BaseFamily/variant"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// CipT 结构体
type CipT struct {
	Text     string // 待处理文本
	Encoding string // 编码方式，默认 UTF-8
	Method   string // 编码/解码方法
}

// NewCipT 构造函数
func NewCipT(text, method string) *CipT {
	return &CipT{
		Text:     text,
		Encoding: "UTF-8", // 设置默认值
		Method:   method,
	}
}

// 定义编码和解码函数类型
type codecFunc func(data []byte) ([]byte, error)

// 编码和解码映射表
var encoderFunc = map[string]codecFunc{
	"ASCII85":  ascii85.StdCodec.Encode,
	"Base4":    base4.StdCodec.Encode,
	"Base8":    base8.StdCodec.Encode,
	"Base16":   base16.StdCodec.Encode,
	"Base24":   base24.StdCodec.Encode,
	"Base32":   base32.StdCodec.Encode,
	"Base36":   base36.StdCodec.Encode,
	"Base45":   base45.StdCodec.Encode,
	"Base58":   base58.StdCodec.Encode,
	"Base64":   base64.StdCodec.Encode,
	"Base85":   base85.StdCodec.Encode,
	"Base91":   base91.StdCodec.Encode,
	"Base92":   base92.StdCodec.Encode,
	"Base100":  base100.StdCodec.Encode,
	"UUEncode": BaseVariant.UUEncode.Encode,
	"XXEncode": BaseVariant.XXEncode.Encode,
	"URL": func(data []byte) ([]byte, error) {
		return []byte(url.QueryEscape(string(data))), nil
	},
}

var decoderFunc = map[string]codecFunc{
	"ASCII85":  ascii85.StdCodec.Decode,
	"Base4":    base4.StdCodec.Decode,
	"Base8":    base8.StdCodec.Decode,
	"Base16":   base16.StdCodec.Decode,
	"Base24":   base24.StdCodec.Decode,
	"Base32":   base32.StdCodec.Decode,
	"Base36":   base36.StdCodec.Decode,
	"Base45":   base45.StdCodec.Decode,
	"Base58":   base58.StdCodec.Decode,
	"Base64":   base64.StdCodec.Decode,
	"Base85":   base85.StdCodec.Decode,
	"Base91":   base91.StdCodec.Decode,
	"Base92":   base92.StdCodec.Decode,
	"Base100":  base100.StdCodec.Decode,
	"UUEncode": BaseVariant.UUEncode.Decode,
	"XXEncode": BaseVariant.XXEncode.Decode,
	"URL": func(data []byte) ([]byte, error) {
		res, err := url.QueryUnescape(string(data))
		return []byte(res), err
	},
}

// 编码转换映射表
var encodings = map[string]struct {
	ToUTF8   codecFunc
	FromUTF8 codecFunc
}{
	"UTF-8":    {identity, identity},
	"GBK":      {decodeTransform(simplifiedchinese.GBK.NewDecoder()), encodeTransform(simplifiedchinese.GBK.NewEncoder())},
	"GB2312":   {decodeTransform(simplifiedchinese.GB18030.NewDecoder()), encodeTransform(simplifiedchinese.GB18030.NewEncoder())},
	"GB18030":  {decodeTransform(simplifiedchinese.GB18030.NewDecoder()), encodeTransform(simplifiedchinese.GB18030.NewEncoder())},
	"HZGB2312": {decodeTransform(simplifiedchinese.HZGB2312.NewDecoder()), encodeTransform(simplifiedchinese.HZGB2312.NewEncoder())},
}

// identity 是一个不做任何转换的函数
func identity(data []byte) ([]byte, error) {
	return data, nil
}

// decodeTransform 创建解码转换函数
func decodeTransform(transformer *encoding.Decoder) codecFunc {
	return func(data []byte) ([]byte, error) {
		reader := transform.NewReader(bytes.NewReader(data), *transformer)
		return io.ReadAll(reader)
	}
}

// encodeTransform 创建编码转换函数
func encodeTransform(transformer *encoding.Encoder) codecFunc {
	return func(data []byte) ([]byte, error) {
		reader := transform.NewReader(bytes.NewReader(data), *transformer)
		return io.ReadAll(reader)
	}
}

// 统一处理逻辑
func (t *CipT) process(isEncode bool) (string, error) {
	// 获取转换函数
	Encoding := encodings[t.Encoding]
	if Encoding.ToUTF8 == nil || Encoding.FromUTF8 == nil {
		return "", errors.New("unknown encoding: " + t.Encoding)
	}

	// 编码或解码主逻辑
	var codec codecFunc
	if isEncode {
		codec = encoderFunc[t.Method]
	} else {
		codec = decoderFunc[t.Method]
	}
	if codec == nil {
		return "", errors.New("unknown method: " + t.Method)
	}

	// 转换文本
	data, err := []byte(t.Text), error(nil)
	if isEncode {
		data, err = Encoding.FromUTF8(data)
		if err != nil {
			return "", err
		}
		data, err = codec(data)
	} else {
		data, err = codec(data)
		if err != nil {
			return "", err
		}
		data, err = Encoding.ToUTF8(data)
	}
	return string(data), err
}

// Encode 编码方法
func (t *CipT) Encode() (string, error) {
	return t.process(true)
}

// Decode 解码方法
func (t *CipT) Decode() (string, error) {
	return t.process(false)
}
