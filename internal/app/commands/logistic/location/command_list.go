package location

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/location"
)

const defaultListLimit = 10

func (c *LogisticLocationCommander) List(inputMsg *tgbotapi.Message) {
	c.listLimited(inputMsg, 0, defaultListLimit)
}

func (c *LogisticLocationCommander) listLimited(inputMsg *tgbotapi.Message, cursor uint64, limit uint64) {
	locations, err := c.locationService.List(cursor, limit)

	if err == location.ErrEmptyList {
		log.Printf("%s.List: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "There's nothing to list", logPrefix+".List")
		return
	} else if err != nil && err != location.ErrEndOfList {
		log.Printf("%s.List: failed to list locations: %v", logPrefix, err)
		c.sendMessage(inputMsg.Chat.ID, "Error occurred", logPrefix+".List")
		return
	}

	var sb strings.Builder

	for _, l := range locations {
		sb.WriteString(l.String())
		sb.WriteString("\n")
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, sb.String())

	if err != location.ErrEndOfList {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: cursor + limit,
			Limit:  defaultListLimit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "logistic",
			Subdomain:    "location",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Next %d", defaultListLimit), callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LogisticLocationCommander.List: error sending reply message to chat - %v", err)
	}
}
