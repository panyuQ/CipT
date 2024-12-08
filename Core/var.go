package Core

import "CipT/Core/NoKey"

var (
	AllNoKeyEncoder = NoKey.GetMethods(true)
	AllNoKeyDecoder = NoKey.GetMethods(false)
)
