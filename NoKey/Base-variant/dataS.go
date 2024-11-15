package Base_variant

import (
	"errors"
)

var (
	// 编码 字符集
	charsetXXEncode = []byte("+-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	charsetUUEncode = []byte(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_")

	// 错误
	conversionFailed = errors.New("conversion failed")
)

func hasBlockEncode(data []byte, charset []byte, blockSize int) ([]byte, error) {
	var encode []byte

	length := len(data)
	var n int
	if n = 3 - length%3; n < 3 {
		data = append(data, make([]byte, n)...)
	}
	var d []byte
	var i int
	num := length
	for i = 0; i < length+n; i += 3 {
		if i > 0 && i%blockSize == 0 {
			encode = append(encode, charset[blockSize])
			encode = append(encode, d...)
			encode = append(encode, '\n')
			d = []byte{}
			num -= blockSize
		}
		d = append(d, charset[data[i]>>2])                           // 第一组：data[0] 的高 6 位
		d = append(d, charset[((data[i]&0x03)<<4)|(data[i+1]>>4)])   // 第二组：data[0] 的低 2 位和 data[1] 的高 4 位
		d = append(d, charset[((data[i+1]&0x0F)<<2)|(data[i+2]>>6)]) // 第三组：data[1] 的低 4 位和 data[2] 的高 2 位
		d = append(d, charset[data[i+2]&0x3F])                       // 第四组：data[2] 的低 6 位

	}
	if num > 0 {
		encode = append(encode, charset[num])
		encode = append(encode, d...)
	}

	n = len(encode)
	res := make([]byte, n)
	copy(res, encode)
	if n > 0 {
		return res, nil
	} else {
		return nil, conversionFailed
	}
}

func hasBlockDecode(data []byte, charset []byte) ([]byte, error) {

	// 构建一个映射表，用于快速查找字符在字符集中的位置
	charIndex := make(map[byte]int)
	for i, c := range charset {
		charIndex[c] = i
	}

	var decode []byte
	length := len(data)
	var num int
	// 解码
	for i := 0; i < length; i += num {
		if data[i] == '\r' {
			i++
		}
		if data[i] == '\n' {
			i++
		}
		num = charIndex[data[i]]
		if num%3 != 0 {
			num += 3 - num%3
		}
		num = num / 3 * 4
		i++
		for j := i; j < i+num; j += 4 {
			b0 := charIndex[data[j+0]]
			b1 := charIndex[data[j+1]]
			b2 := charIndex[data[j+2]]
			b3 := charIndex[data[j+3]]

			byte0 := byte((b0 << 2) | (b1 >> 4))
			byte1 := byte((b1 << 4) | (b2 >> 2))
			byte2 := byte((b2 << 6) | b3)

			decode = append(decode, byte0, byte1, byte2)
		}
	}
	n := len(decode)
	res := make([]byte, n)
	copy(res, decode)
	if n > 0 {
		return res, nil
	} else {
		return nil, conversionFailed
	}
}
