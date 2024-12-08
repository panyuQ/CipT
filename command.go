package main

import (
	"CipT/Proc"
	"flag"
	"fmt"
	"os"
)

var VERSION = "1.0.0"

var (
	FlagBool = map[string]bool{
		"window": *flag.Bool("window", false, "Windowed mode"),
		"web":    *flag.Bool("web", false, "Web application mode"),
		"encode": *flag.Bool("encode", false, "Encode mode"),
	}

	FlagString = map[string]string{
		"method": *flag.String("method", "", "En/Decode method"),
		"output": *flag.String("output", "", "File path (默认)"),
	}

	FlagInt = map[string]int{
		"page": *flag.Int("page", 0, "Page number"),
	}
)

func changeHelp() {
	flag.CommandLine.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
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
	proc := Proc.NewProc(FlagString["method"])
	proc.Run(FlagInt["page"])

}
