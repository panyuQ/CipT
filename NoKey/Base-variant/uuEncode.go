package Base_variant

func UUEncode(data []byte) ([]byte, error) {
	return hasBlockEncode(data, charsetUUEncode, 45)
}

func UUDecode(data []byte) ([]byte, error) {
	return hasBlockDecode(data, charsetUUEncode)
}
