package Proc

type Proc struct {
	content string
	file    bool
	encode  bool
	method  string
}

func NewProc(content string, file bool, encode bool, method string) *Proc {
	return &Proc{content, file, encode, method}
}

func (p *Proc) GetFunc() {

}
