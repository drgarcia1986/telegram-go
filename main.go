package main

import (
	"fmt"
	"os"

	"github.com/drgarcia1986/telegram-go/server"
)

func greeting(msg string, mv map[string]string) (string, error) {
	return fmt.Sprintf("Hello %s nice to meet you", mv["name"]), nil
}

func main() {
	s := server.New(os.Getenv("TOKEN"))
	s.HandleFunc("Hi I'm (?P<name>\\w+)", greeting)

	fmt.Println("Start Telegram BOT")
	s.Run()
}
