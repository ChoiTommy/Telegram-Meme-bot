package main
/* Meme bot
 * Source: https://github.com/go-telegram-bot-api/telegram-bot-api
 */
import (
	"log"
	"strings"
	"github.com/go-telegram-bot-api/telegram-bot-api"
  "os"
  "net/http"
  "math/rand"
  "time"
  "fmt"
)

func main_bot() {
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

	list := [8]string{"偷", "練", "做", "睇", "食", "瞓", "玩", "問"}
	messages := [4]string{"%s完又%s", "日%s夜%s", "一%s再%s", "死%s爛%s"}

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		i := 0
		match := false
		for ((i <= 7) && !match) {
			match = strings.Contains(update.Message.Text, list[i])
			i++
		}

		if match {
			pending := ""
			for n := 0; n <= 3; n++ {
				pending = pending + strings.ReplaceAll(messages[n], "%s", list[i-1]) + "\n"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, pending)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

    if strings.Contains(strings.ToUpper(update.Message.Text), "DSE"){
      file := r1.Intn(12) + 1
			log.Printf("File: %d.webp", file)
			path := fmt.Sprintf("stickers/%d.webp", file)
			log.Printf(path)
			msgReply := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, path)
			msgReply.ReplyToMessageID = update.Message.MessageID
			bot.Send(msgReply)
			log.Printf("Sticker %d sent", file)
    }
	}
}
