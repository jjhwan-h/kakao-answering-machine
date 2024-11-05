package main

import (
	"chatbot/cmd"
	"time"
)

func main() {
	cmd.Execute()
	time.Sleep(1 * time.Second)
}
