package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tg "github.com/Syfaro/telegram-bot-api"
)

var (
	bot    *tg.BotAPI
	config Config
)

type FoodRecord struct {
	Results []struct {
		Name string
	}
}

func main() {
	config = LoadConfigFile("./config.json")
	var err error
	bot, err = tg.NewBotAPI(config.Token)
	resp := &http.Response{}
	googleMapsURL := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?"

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("123 %s", bot.Self.UserName)

	var ucfg tg.UpdateConfig
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
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
			} else if message.Location != nil {
				messageLocation := fmt.Sprintf("Location -->> %.6f %.6f",
					message.Location.Latitude, message.Location.Longitude)
				log.Printf(messageLocation)
				urlGet := googleMapsURL +
					"location=" + fmt.Sprintf("%.6f,%.6f", message.Location.Latitude,
					message.Location.Longitude) +
					"&radius=500" +
					"&type=restaurant" +
					"&key=" + config.GoogleApiKey

				resp, _ = http.Get(string(urlGet))
				data, _ := ioutil.ReadAll(resp.Body)
				var record FoodRecord
				err = json.Unmarshal(data, &record)
				if err != nil {
					log.Println(err)
				}
				for place := range record.Results {
					sendMessage(chatID, string(place), nil)
				}
				log.Println("asdasd", record.Results)
			}
		}
	}
}
