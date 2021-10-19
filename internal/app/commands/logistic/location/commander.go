package location

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/logistic/location"
	"log"
)

const logPrefix = "LogisticLocationCommander"

type LocationCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type LogisticLocationCommander struct {
	bot             *tgbotapi.BotAPI
	locationService service.LocationService
}

func NewLocationCommander(bot *tgbotapi.BotAPI, service service.LocationService) *LogisticLocationCommander {
	return &LogisticLocationCommander{
		bot:             bot,
		locationService: service,
	}
}

func (c *LogisticLocationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("%s.HandleCallback: unknown callback name: %s", logPrefix, callbackPath.CallbackName)
	}
}

func (c *LogisticLocationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *LogisticLocationCommander) sendMessage(chatID int64, text string, logPrefix string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("%s: error sending reply message to chat - %v", logPrefix, err)
	}
}
