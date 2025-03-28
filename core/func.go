package core

import (
	"CipT/core/BaseFamily/codec/ascii85"
	"CipT/core/BaseFamily/codec/base100"
	"CipT/core/BaseFamily/codec/base16"
	"CipT/core/BaseFamily/codec/base24"
	"CipT/core/BaseFamily/codec/base32"
	"CipT/core/BaseFamily/codec/base36"
	"CipT/core/BaseFamily/codec/base4"
	"CipT/core/BaseFamily/codec/base45"
	"CipT/core/BaseFamily/codec/base58"
	"CipT/core/BaseFamily/codec/base64"
	"CipT/core/BaseFamily/codec/base8"
	"CipT/core/BaseFamily/codec/base85"
	"CipT/core/BaseFamily/codec/base91"
	"CipT/core/BaseFamily/codec/base92"
	"CipT/core/BaseFamily/variant/HasBlock"
	"net/url"
)

// 编码和解码映射表
var encoderFunc = map[string]codecFunc{
	"ASCII85":   ascii85.StdCodec.Encode,
	"Base4":     base4.StdCodec.Encode,
	"Base8":     base8.StdCodec.Encode,
	"Base16":    base16.StdCodec.Encode,
	"Base24":    base24.StdCodec.Encode,
	"Base32":    base32.StdCodec.Encode,
	"Base36":    base36.StdCodec.Encode,
	"Base45":    base45.StdCodec.Encode,
	"Base58":    base58.StdCodec.Encode,
	"Base64":    base64.StdCodec.Encode,
	"Base64Url": base64.UrlCodec.Encode,
	"Base85":    base85.StdCodec.Encode,
	"Base91":    base91.StdCodec.Encode,
	"Base92":    base92.StdCodec.Encode,
	"Base100":   base100.StdCodec.Encode,
	"UUEncode":  HasBlock.UUEncode.Encode,
	"XXEncode":  HasBlock.XXEncode.Encode,
	"URL": func(data []byte) ([]byte, error) {
		return []byte(url.QueryEscape(string(data))), nil
	},
}

var decoderFunc = map[string]codecFunc{
	"ASCII85":   ascii85.StdCodec.Decode,
	"Base4":     base4.StdCodec.Decode,
	"Base8":     base8.StdCodec.Decode,
	"Base16":    base16.StdCodec.Decode,
	"Base24":    base24.StdCodec.Decode,
	"Base32":    base32.StdCodec.Decode,
	"Base36":    base36.StdCodec.Decode,
	"Base45":    base45.StdCodec.Decode,
	"Base58":    base58.StdCodec.Decode,
	"Base64":    base64.StdCodec.Decode,
	"Base64Url": base64.UrlCodec.Decode,
	"Base85":    base85.StdCodec.Decode,
	"Base91":    base91.StdCodec.Decode,
	"Base92":    base92.StdCodec.Decode,
	"Base100":   base100.StdCodec.Decode,
	"UUEncode":  HasBlock.UUEncode.Decode,
	"XXEncode":  HasBlock.XXEncode.Decode,
	"URL": func(data []byte) ([]byte, error) {
		res, err := url.QueryUnescape(string(data))
		return []byte(res), err
	},
}

func GetMethods(encode bool) []string {
	var result []string
	if encode {
		for method := range encoderFunc {
			result = append(result, method)
		}

	} else {
		for method := range decoderFunc {
			result = append(result, method)
		}
	}
	return result
}
