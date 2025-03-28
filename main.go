package main

import (
	"CipT/proc/Web"
	"CipT/proc/Window"
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
