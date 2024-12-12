package main

import (
	"CipT/Proc/Web"
	"CipT/Proc/Window"
)

func main() {
	Command()
	if *FlagBool["window"] {
		Window.Run()

	}
	if *FlagBool["web"] {
		Web.Run()
	}
}
