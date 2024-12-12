package Proc

const (
	Base64                 = "^[A-Za-z0-9+/]+={0,2}$"
	Base64UrlSafe          = "^[A-Za-z0-9_-]+={0,2}$"
	Base64NoPadding        = "^[A-Za-z0-9+/]+$"
	Base64UrlSafeNoPadding = "^[A-Za-z0-9_-]+$"
	Hex                    = "^[0-9a-fA-F]+$"
	MD5                    = "^[a-fA-F0-9]{32}$"
	SHA1                   = "^[a-fA-F0-9]{40}$"
)
