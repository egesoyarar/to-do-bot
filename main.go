package main

import (
	"fmt"
	"os"
	chat "to-do-bot/chat_bot"
)

var (
	token = os.Getenv("TOKEN")
)

func main() {
	_, err := chat.NewTelegramBot(token)
	if err != nil {
		fmt.Printf("Telegram client cannot be created due to %v", err)
	}
}
