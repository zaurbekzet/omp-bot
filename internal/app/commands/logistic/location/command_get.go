package location

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticLocationCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("%s.Get: wrong arguments: %s", logPrefix, args)
		c.sendMessage(
			inputMsg.Chat.ID,
			"Wrong argument. Use this format: `/get__logistic__location <id>`",
			logPrefix+".Get",
		)
		return
	}

	location, err := c.locationService.Describe(uint64(id))
	if err != nil {
		log.Printf("%s.Get: failed to get location: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "Requested location does not exist", logPrefix+".Get")
		return
	}

	c.sendMessage(inputMsg.Chat.ID, location.String(), logPrefix+".Get")
}
