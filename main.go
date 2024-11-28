package main

func main() {
	Command()
	if FlagBool["window"] {
		Window()

	}
	if FlagBool["web"] {
		Web()
	}

	task()
}
