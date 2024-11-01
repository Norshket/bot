package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("recoverd from panic : %v", r)
		}
	}()

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		fmt.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)

	if err != nil {
		fmt.Printf("Fail to the product with idx %d : %v", args, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
