package chat_bot

type IChatClient interface {
	SendMessage(sendMsg string, chatId int64) error
	ReceiveMessage() (receivedMsg string)
}

type ChatClient struct {
	token string
}

func NewChatClient(token string) IChatClient {
	return &ChatClient{token: token}
}

func (c *ChatClient) SendMessage(sendMsg string, chatId int64) (err error) {
	return err
}

func (c *ChatClient) ReceiveMessage() (receivedMsg string) {
	return receivedMsg
}
