package variant

import (
	"bytes"
	"errors"
)

const (
	xxEncoder = "+-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	uuEncoder = "`!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_"
)

type hasBlockEncoder struct {
	charset   []byte
	blockSize int
}

var (
	XXEncode = hasBlockEncoder{
		charset:   []byte(xxEncoder),
		blockSize: 45,
	}
	UUEncode = hasBlockEncoder{
		charset:   []byte(uuEncoder),
		blockSize: 45,
	}
	conversionFailed = errors.New("has block encoder conversion failed")
)

// Encode 将数据编码为块格式
func (config *hasBlockEncoder) Encode(data []byte) ([]byte, error) {
	length := len(data)
	data = append(data, make([]byte, 3-length%3)...)

	var buffer bytes.Buffer

	for i := 0; i < length; i += 3 {
		if i%config.blockSize == 0 {
			if i > 0 {
				buffer.WriteByte('\n')
			}
			if i+config.blockSize > length { // 如果是最后一行
				buffer.WriteByte(config.charset[length-i])
			} else {
				buffer.WriteByte(config.charset[config.blockSize])
			}

		}
		buffer.WriteByte(config.charset[data[i]>>2])
		buffer.WriteByte(config.charset[((data[i]&0x03)<<4)|(data[i+1]>>4)])
		buffer.WriteByte(config.charset[((data[i+1]&0x0F)<<2)|(data[i+2]>>6)])
		buffer.WriteByte(config.charset[data[i+2]&0x3F])

	}

	// 在结尾处添加块的长度字符
	return buffer.Bytes(), nil
}

// Decode 将块格式的数据解码为原始数据
func (config *hasBlockEncoder) Decode(data []byte) ([]byte, error) {
	charIndex := make(map[byte]int, len(config.charset))
	for i, c := range config.charset {
		charIndex[c] = i
	}

	var decoded bytes.Buffer
	lines := bytes.Split(data, []byte{'\n'})

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		//num := charIndex[line[0]]
		for i := 1; i < len(line); i += 4 {
			if i+3 >= len(line) {
				break
			}
			b0 := charIndex[line[i+0]]
			b1 := charIndex[line[i+1]]
			b2 := charIndex[line[i+2]]
			b3 := charIndex[line[i+3]]

			decoded.WriteByte(byte((b0 << 2) | (b1 >> 4)))
			decoded.WriteByte(byte((b1 << 4) | (b2 >> 2)))
			decoded.WriteByte(byte((b2 << 6) | b3))
		}
	}

	result := decoded.Bytes()
	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}
	return result, nil
}
