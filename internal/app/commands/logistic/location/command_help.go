package location

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticLocationCommander) Help(inputMsg *tgbotapi.Message) {
	msgText := `/help__logistic__location — print list of commands
/get__logistic__location <id> — get location
/list__logistic__location — get list of locations
/delete__logistic__location <id> — delete existing location
/new__logistic__location {"title": "<title>", "latitude": <latitude>, "longitude": <longitude>} — create new location
/edit__logistic__location {"id": <id>, "title": "<new_title>", "latitude": <new_latitude>, "longitude": <new_longitude>} — edit location`

	c.sendMessage(inputMsg.Chat.ID, msgText, logPrefix+".Help")
}
