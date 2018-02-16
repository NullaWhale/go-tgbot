package main

import (
	"encoding/json"
	tg "github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

type Config struct {
	Token string `json:"Token"`
}

func main() {
	config := LoadConfigFile("./config.json")
	log.Println(config.Token)
	bot, err := tg.NewBotAPI(config.Token)

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
		if update.Message != nil {
			handleMessage(update, bot)
		}

		if update.CallbackQuery != nil {
			handleInlineKeyboardAnswer(update, bot)
		}
	}
}

func handleMessage(update tg.Update, bot *tg.BotAPI) {
	UserName := update.Message.From.UserName
	ChatID := update.Message.Chat.ID
	Message := update.Message

	switch Message.Command() {
	case "start":
		markup := tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Yes", "befat_yes")),
			tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("Of cause!", "befat_ofcause")),
		)

		helloMessage := "Hello, " + UserName + ". Do you want to be fat?"
		msg := tg.NewMessage(int64(ChatID), helloMessage)
		msg.ReplyMarkup = markup
		bot.Send(msg)
	}
}

func handleInlineKeyboardAnswer(update tg.Update, bot *tg.BotAPI) {
	if update.CallbackQuery.Data == "befat_yes" {
		msg := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, "So, that's grate. Stay on your way.")
		bot.Send(msg)
	} else if update.CallbackQuery.Data == "befat_ofcause" {
		msg := tg.NewMessage(update.CallbackQuery.Message.Chat.ID, "Perfect! I like this guy.")
		bot.Send(msg)
	}
}

func LoadConfigFile(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
