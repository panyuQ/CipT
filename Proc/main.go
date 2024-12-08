package Proc

type Proc struct {
	input  Input
	method string
}

func NewProc(method string) *Proc {
	return &Proc{input: *NewInput(), method: method}
}

func (p *Proc) Run(page int) {
	switch p.method {
	case "":

		break
	}
}
