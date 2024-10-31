package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "The all products \n"

	products := c.productService.List()

	for _, product := range products {
		outputMsg += product.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
