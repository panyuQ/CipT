package core

type CipTInterface interface {
	Encode(text ...string) ([]string, error)
	Decode(encodedText ...string) ([]string, error)
}
