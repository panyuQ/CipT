package main

import (
	"CipT/Proc"
	"flag"
	"fmt"
	"os"
)

var VERSION = "1.0.0"

var (
	FlagBool = map[string]*bool{
		"window": flag.Bool("window", false, "Windowed mode"),
		"web":    flag.Bool("web", false, "Web application mode"),
		"encode": flag.Bool("encode", false, "Encode mode"),
		// 只识别密文类型
		"onlyIdentify": flag.Bool("onlyIdentify", false, "Only recognize types (when encode is false)"),
	}

	FlagString = map[string]*string{
		"output": flag.String("output", "./out/", "output directory path"),
		"method": flag.String("method", "", "method name"),
		"key":    flag.String("key", "", "decryption key"),
		"other":  flag.String("other", "", "other parameters"),
	}

	FlagInt = map[string]*int{
		"pageSize":          flag.Int("pageSize", 2000, "Page number"),
		"workers":           flag.Int("workers", 0, "Number of workers ( default number of CPU cores )"),
		"tasksBufferSize":   flag.Int("tasksBufferSize", 0, "Buffer size of tasks ( default workers * 100 )"),
		"resultsBufferSize": flag.Int("resultsBufferSize", 0, "Buffer size of results ( default tasksBufferSize * 100 )"),
	}
)

func changeHelp() {
	flag.CommandLine.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s <method> [text...]:\n", os.Args[0])
		flag.PrintDefaults()
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "\nVersion: %s", VERSION)
	}
}

func Command() {
	changeHelp()
	flag.Parse()
	if flag.NArg() < 1 {
		return
	}
	proc := Proc.NewProc(flag.Args(), *FlagString["method"], *FlagBool["encode"], *FlagString["key"], *FlagString["other"])
	proc.Config.PageSize = *FlagInt["pageSize"]
	proc.Config.Workers = *FlagInt["workers"]
	proc.Config.TasksBufferSize = *FlagInt["tasksBufferSize"]
	proc.Config.ResultsBufferSize = *FlagInt["resultsBufferSize"]
	proc.IdentifyMethod()
	if !*FlagBool["onlyIdentify"] {
		proc.Run()
	}
	proc.Output(*FlagString["output"])
}
