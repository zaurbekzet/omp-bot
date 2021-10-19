package location

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
	"log"
)

func (c *LogisticLocationCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	var location logistic.Location

	if err := json.Unmarshal([]byte(args), &location); err != nil {
		log.Printf("%s.Edit: wrong arguments: %s", logPrefix, args)
		c.sendMessage(
			inputMsg.Chat.ID,
			`Wrong arguments. Use this format:
/edit__logistic__location {"id": <id>, "title": "<new_title>", "latitude": <new_latitude>, "longitude": <new_longitude>}`,
			logPrefix+"Edit",
		)
		return
	}

	if err := c.locationService.Update(location.ID, location); err != nil {
		log.Printf("%s.Get: failed to update location: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "Specified location does not exist", logPrefix+".Edit")
		return
	}

	c.sendMessage(inputMsg.Chat.ID, "Location successfully updated", logPrefix+".Edit")
}
