package core

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// 定义编码和解码函数类型
type codecFunc func(data []byte) ([]byte, error)

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

// identity 是一个不做任何转换的函数
func identity(data []byte) ([]byte, error) {
	return data, nil
}
