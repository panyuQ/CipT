package Base_variant

func XXEncode(data []byte) ([]byte, error) {
	return hasBlockEncode(data, charsetXXEncode, 45)
}

func XXDecode(data []byte) ([]byte, error) {
	return hasBlockDecode(data, charsetXXEncode)
}
