package logistic

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	locationCmd "github.com/ozonmp/omp-bot/internal/app/commands/logistic/location"
	"github.com/ozonmp/omp-bot/internal/app/path"
	locationSrv "github.com/ozonmp/omp-bot/internal/service/logistic/location"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot               *tgbotapi.BotAPI
	locationCommander Commander
}

func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot: bot,
		// locationCommander
		locationCommander: locationCmd.NewLocationCommander(bot, locationSrv.NewDummyLocationService()),
	}
}

func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "location":
		c.locationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "location":
		c.locationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
