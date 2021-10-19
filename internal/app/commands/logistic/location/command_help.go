package location

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LogisticLocationCommander) Help(inputMsg *tgbotapi.Message) {
	msgText := `/help__logistic__location — print list of commands
/get__logistic__location — get location
/list__logistic__location — get list of locations
/delete__logistic__location — delete existing location
/new__logistic__location — create new location
/edit__logistic__location — edit location`

	c.sendMessage(inputMsg.Chat.ID, msgText, logPrefix+".Help")
}
