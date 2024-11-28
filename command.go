package main

import (
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
		"decode": *flag.Bool("decode", true, "Decode mode"),
	}

	FlagString = map[string]string{
		"filepath": *flag.String("filepath", "", "File path (for batch processing)"),
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
}
