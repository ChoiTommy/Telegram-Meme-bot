package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func configure_webhook(bot tgbotapi.BotAPI) tgbotapi.UpdatesChannel{
  _, err := bot.SetWebhook(tgbotapi.NewWebhook("https://Telegram-bot--kaitojjj.repl.co:443/"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:8443", nil)
  return updates
}

/*func main_webhook() {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://Telegram-bot--kaitojjj.repl.co:443/"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}*/
