package chat_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramClient struct {
	bot   *tgbotapi.BotAPI
	token string
}

func NewTelegramClient(token string, chatId int64) (IChatClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	client := &TelegramClient{bot: bot, token: token}
	return client, err
}

func (tc *TelegramClient) SendMessage(sendMsg string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, sendMsg)
	_, err := tc.bot.Send(msg)
	return err
}

func (tc *TelegramClient) ReceiveMessage() (receivedMsg string) {
	return receivedMsg
}
