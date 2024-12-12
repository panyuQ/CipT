package Proc

import (
	"CipT/Core"
	"CipT/Core/NoKey"
	"CipT/Logger"
	"CipT/Task"
)

type Proc struct {
	input  Input
	method string
	encode bool
	key    string
	other  string
	tasks  Task.WorkPool
	Config Config
}

func NewProc(input []string, method string, encode bool, key string, other string) *Proc {
	return &Proc{input: *NewInput(input), method: method, encode: encode, key: key, other: other, Config: *NewConfig()}
}

// IdentifyMethod 识别方法
func (p *Proc) IdentifyMethod() {
	var allMethods []string
	if p.encode {
		allMethods = append(allMethods, Core.AllNoKeyEncoder...)
		// 添加有密钥的编码
	} else {
		allMethods = append(allMethods, Core.AllNoKeyDecoder...)
		// 添加有密钥的解码
	}
	for _, str := range allMethods {
		if str == p.method {
			return
		}
	}

	// 识别代码

	// 识别识别不出就返回空
	return
}

func (p *Proc) Run() {
	if p.method == "" {
		Logger.GLogger.Error.Println(methodNoExist, p.method)
		return
	} else {
		if p.key == "" {
			p.runNoKey()
		} else {

		}
	}
}

func (p *Proc) Output(filepath string) {
	p.tasks.Output(filepath)
}

func (p *Proc) runNoKey() {
	cipT := NoKey.NewCipT(p.method)

	p.tasks = *Task.NewWorkPool(p.Config.Workers, p.Config.TasksBufferSize, p.Config.ResultsBufferSize, cipT.Encode, Logger.GLogger)
	p.tasks.Start()
	var i int
	count := p.input.GetAllContentCount()
	for i = range count / p.Config.PageSize {
		p.tasks.AddTask(Task.NewTask(i, p.input.GetContent(i, p.Config.PageSize)))
		count -= p.Config.PageSize
	}
	p.tasks.AddTask(Task.NewTask(i, p.input.GetContent(i, p.Config.PageSize))) // 添加剩余任务
	p.tasks.Stop(false)
}
