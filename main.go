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
	config := GetConfig()

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotKey)
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
			Key: config.FinhubApiKey,
		})

		quote, _, err := finhubClient.Quote(auth, "WIX")

		if err != nil{
			fmt.Printf("Oi yooo.")
		}

		fmt.Printf("%+v\n", quote)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		price := strconv.FormatFloat(float64(quote.Pc), 'f', 4, 64)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, price)
		msg.ReplyToMessageID = update.Message.MessageID

		_, _ = bot.Send(msg)
	}
}
