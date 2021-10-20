package location

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticLocationCommander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	c.sendMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text, logPrefix+".Default")
}
