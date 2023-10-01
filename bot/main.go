package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	// Get token from the environment variable
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	// This MUST be an HTTPS URL for Telegram to accept it
	webAppUrl := os.Getenv("URL")
	if webAppUrl == "" {
		panic("URL environment variable is empty")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client:             http.Client{},
			UseTestEnvironment: true,
			DefaultRequestOpts: nil,
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher to handle updates in a simple manner.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dispatcher := updater.Dispatcher

	// /start command to introduce the bot and send the URL
	dispatcher.AddHandler(handlers.NewCommand("start", func(b *gotgbot.Bot, ctx *ext.Context) error {
		return start(b, ctx, webAppUrl)
	}))

	// Start receiving (and handling) updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

// start introduces the bot.
func start(b *gotgbot.Bot, ctx *ext.Context, webAppUrl string) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s.\nYou can use me to run a (very) simple telegram webapp demo!", b.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Press me", WebApp: &gotgbot.WebAppInfo{Url: webAppUrl}},
			}},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
