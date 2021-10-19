package location

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticLocationCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("%s.Delete: wrong arguments: %s", logPrefix, args)
		c.sendMessage(
			inputMsg.Chat.ID,
			"Wrong argument. Use this format: `/delete__logistic__location <locationID>`",
			logPrefix+".Delete",
		)
		return
	}

	_, err = c.locationService.Remove(uint64(id))
	if err != nil {
		log.Printf("%s.Delete: failed to delete location: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "Specified location does not exist", logPrefix+".Delete")
		return
	}

	c.sendMessage(inputMsg.Chat.ID, "Location successfully deleted", logPrefix+".Delete")
}
