package main

import (
	"CipT/NoKey"
	"testing"
)

var (
	tests = []struct {
		Plaintext  string            // Before 编码前/解码后
		Ciphertext map[string]string // After 编码后/解码前
		Method     string            // Method 方法
	}{
		{
			Plaintext: "Hello, World!",
			Ciphertext: map[string]string{
				"UTF-8": "SGVsbG8sIFdvcmxkIQ==",
			},
			Method: "Base64",
		},
		{
			Plaintext: "你好，世界！",
			Ciphertext: map[string]string{
				"UTF-8": "5L2g5aW977yM5LiW55WM77yB",
			},
			Method: "Base64",
		},
	}
)

func TestAll(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.Method, func(t *testing.T) {
			cipT := NoKey.CipT{Method: tt.Method}
			for _, encoding := range NoKey.AllEncodings {
				cipT.Encoding = encoding // 设置编码

				cipT.Text = tt.Plaintext      // 设置 明文
				encoded, err := cipT.Encode() // 编码
				target := tt.Ciphertext[encoding]
				if err == nil && (target == "" || encoded == tt.Ciphertext[encoding]) {
					t.Logf("(%s) Successfully encoded using %s (Got: %s)\n", tt.Method, encoding, encoded)
				} else {
					t.Errorf("(%s) Failed encoded using %s (Got: %s, Want: %s)\n%v", tt.Method, encoding, encoded, tt.Ciphertext, err)
				}

				if tt.Ciphertext[encoding] == "" {
					cipT.Text = encoded
				} else {
					cipT.Text = tt.Ciphertext[encoding]
				}
				decoded, err := cipT.Decode()
				if err == nil && decoded == tt.Plaintext {
					t.Logf("(%s) Successfully decoded using %s (Got: %s)\n", tt.Method, encoding, decoded)
				} else {
					t.Errorf("(%s) Failed encoded using %s (Got: %s, Want: %s)\n%v", tt.Method, encoding, decoded, tt.Plaintext, err)
				}

			}
		})
	}

}
