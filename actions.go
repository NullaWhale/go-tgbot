package main

import (
	tg "github.com/Syfaro/telegram-bot-api"
	"log"
	"reflect"
)

func sendMessage(chatID int64, message string, keyboard interface{}) {
	msg := tg.NewMessage(chatID, message)
	typeOfKeyboard := reflect.TypeOf(keyboard)
	log.Println(typeOfKeyboard)
	if typeOfKeyboard == nil {
		msg.ReplyMarkup = tg.ReplyKeyboardRemove{true, false}
	} else {
		switch typeOfKeyboard.String() {
		default:
			msg.ReplyMarkup = tg.ReplyKeyboardRemove{true, false}
			bot.Send(msg)
		case "tgbotapi.InlineKeyboardMarkup":
			msg.ReplyMarkup = &keyboard
			bot.Send(msg)
		}
	}
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
