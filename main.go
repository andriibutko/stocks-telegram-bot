package main

import (
	"context"
	"fmt"
	finhub "github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("telegram-bot-api")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		finhubClient := finhub.NewAPIClient(finhub.NewConfiguration()).DefaultApi
		auth := context.WithValue(context.Background(), finhub.ContextAPIKey, finhub.APIKey{
			Key: "API-KEY",
		})

		quote, _, err := finhubClient.Quote(auth, "WIX")

		if err == nil {
			fmt.Print("Finhub Client uncreated.")
		}

		fmt.Printf("%+v\n", quote)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var price = strconv.FormatFloat(float64(quote.Pc), 'f', 4, 64)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, price)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}