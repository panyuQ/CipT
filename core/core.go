package core

import (
	"errors"
)

// CipT 结构体
type CipT struct {
	Encoding string // 编码方式，默认 UTF-8
	Method   string // 编码/解码方法
}

// NewCipT 构造函数
func NewCipT(method string) *CipT {
	return &CipT{
		Encoding: "UTF-8", // 设置默认值
		Method:   method,
	}
}

// 统一处理逻辑
func (t *CipT) process(isEncode bool, text []string) ([]string, error) {
	// 获取转换函数
	Encoding := encodings[t.Encoding]
	if Encoding.ToUTF8 == nil || Encoding.FromUTF8 == nil {
		return nil, errors.New("unknown encoding: " + t.Encoding)
	}

	// 编码或解码主逻辑
	var codec codecFunc
	if isEncode {
		codec = encoderFunc[t.Method]
	} else {
		codec = decoderFunc[t.Method]
	}
	if codec == nil {
		return nil, errors.New("unknown method: " + t.Method)
	}

	// 转换文本
	result := make([]string, len(text))
	for i, item := range text {
		data := []byte(item)
		var err error
		if isEncode {
			data, err = Encoding.FromUTF8(data)
			if err != nil {
				return nil, err
			}
			data, err = codec(data)
		} else {
			data, err = codec(data)
			if err != nil {
				return nil, err
			}
			data, err = Encoding.ToUTF8(data)
		}
		if err != nil {
			return nil, err
		}
		result[i] = string(data)
	}

	return result, nil
}

// Encode 编码方法
func (t *CipT) Encode(text []string) ([]string, error) {
	return t.process(true, text)
}

// Decode 解码方法
func (t *CipT) Decode(text []string) ([]string, error) {
	return t.process(false, text)
}
