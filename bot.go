package main

import (
	tg "github.com/Syfaro/telegram-bot-api"
	"log"
)

var (
	bot    *tg.BotAPI
	config Config
)

func main() {
	config = LoadConfigFile("./config.json")
	var err error
	bot, err = tg.NewBotAPI(config.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("123 %s", bot.Self.UserName)

	var ucfg tg.UpdateConfig
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		log.Println("-->> ", update)

		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			message := update.CallbackQuery.Message
			handleCallback(data, message)
		} else if update.Message != nil {
			userName := update.Message.From.UserName
			chatID := update.Message.Chat.ID
			message := update.Message
			if message.IsCommand() {
				switch message.Command() {
				case "start":
					markup := tg.NewInlineKeyboardMarkup(
						tg.NewInlineKeyboardRow(
							tg.NewInlineKeyboardButtonData("Yes", "befat_yes"),
							tg.NewInlineKeyboardButtonData("Of cause!", "befat_ofcause"),
						),
					)

					helloMessage := "Hello, " + userName + ". Do you want to be fat?"
					sendMessage(chatID, helloMessage, markup)
				}
			}
		}
	}
}
