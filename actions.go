package main

import (
	"log"
	"reflect"

	tg "github.com/Syfaro/telegram-bot-api"
)

func sendMessage(chatID int64, message string, keyboard interface{}) {
	msg := tg.NewMessage(chatID, message)
	typeOfKeyboard := reflect.TypeOf(keyboard)
	log.Println(typeOfKeyboard)
	if typeOfKeyboard == nil {
		bot.Send(msg)
	} else {
		switch typeOfKeyboard.String() {
		default:
			bot.Send(msg)
		case "tgbotapi.InlineKeyboardMarkup":
			msg.ReplyMarkup = &keyboard
			bot.Send(msg)
		}
	}
}

func sendLocation(chatID int64, latitude float64, longitude float64) {
	msg := tg.NewLocation(chatID, latitude, longitude)
	bot.Send(msg)
}

func handleCallback(data string, message *tg.Message) {
	removeKeyboard := tg.NewEditMessageText(
		message.Chat.ID,
		message.MessageID,
		message.Text)
	bot.Send(removeKeyboard)
	if data == "befat_yes" {
		msg := tg.NewMessage(message.Chat.ID, "So, that's grate. Stay on your way.")
		bot.Send(msg)
	} else if data == "befat_ofcause" {
		msg := tg.NewMessage(message.Chat.ID, "Perfect! I like this guy.")
		bot.Send(msg)
	}
}
