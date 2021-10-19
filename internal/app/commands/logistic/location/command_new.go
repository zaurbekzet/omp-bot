package location

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"log"
)

func (c *LogisticLocationCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	var location logistic.Location

	if err := json.Unmarshal([]byte(args), &location); err != nil {
		log.Printf("%s.New: wrong arguments: %s", logPrefix, args)
		c.sendMessage(
			inputMsg.Chat.ID,
			`Wrong arguments. Use this format:
/new__logistic__location {"title": "<title>", "latitude": <latitude>, "longitude": <longitude>}`,
			logPrefix+"New",
		)
		return
	}

	if _, err := c.locationService.Create(location); err != nil {
		log.Printf("%s.Get: failed to create location: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "Error occurred", logPrefix+".New")
		return
	}

	c.sendMessage(inputMsg.Chat.ID, "Location successfully created", logPrefix+".New")
}
