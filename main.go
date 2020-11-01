package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	finhub "github.com/Finnhub-Stock-API/finnhub-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const welcomeMessage = `Hi, I can tell you how much money you got from WIX stocks.
Use command /price to get current Stock price.`

const stubCommand =
`This feature under development.
You can try to donate money to speed up this process.`

const defaultResponse =
`Try to use commands, I don't speak human language.`

const rejectResponse =
`You are not from WIX, GFY.`

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

		if err != nil {
			fmt.Printf("Oi yooo.")
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		stockPrice := getStockPrice("WIX")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			getCommandResponse(update.Message.Text, stockPrice))

		_, _ = bot.Send(msg)
	}
}

func getCommandResponse(input string, price float32) string {
	var text string

	switch input {
	case "/start":
		text = welcomeMessage
		break
	case "/price":
		text = strconv.FormatFloat(float64(price), 'f', 4, 64)
		break
	case "/subscribe":
		text = stubCommand
	default:
		text = defaultResponse
	}

	return text
}

func getStockPrice(stock string) float32 {
	config := GetConfig()

	finhubClient := finhub.NewAPIClient(finhub.NewConfiguration()).DefaultApi
	auth := context.WithValue(context.Background(), finhub.ContextAPIKey, finhub.APIKey{
		Key: config.FinhubAPIKey,
	})

	quote, _, err := finhubClient.Quote(auth, stock)
	if err != nil {
		log.Printf("%+v\n", quote)
	}

	return quote.Pc
}