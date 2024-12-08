package Proc

import (
	"CipT/Core"
	"CipT/Logger"
)

type Proc struct {
	input  Input
	method string
	encode bool
}

func NewProc(method string) *Proc {
	return &Proc{input: *NewInput(), method: method}
}

func (p *Proc) Run(page int) {
	if p.method == "" {
		Logger.GLogger.Error.Println(methodNoExist, p.method)
		return
	}

}

// identifyMethod 识别方法
func (p *Proc) identifyMethod() string {
	var allMethods []string
	if p.encode {
		allMethods = append(allMethods, Core.AllNoKeyEncoder...)
	} else {
		allMethods = append(allMethods, Core.AllNoKeyDecoder...)
	}
	for _, str := range allMethods {
		if str == p.method {
			return p.method
		}
	}

	// 识别识别不出就返回空
	return ""
}
