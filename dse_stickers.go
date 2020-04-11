package main

import(
  "strings"
  "time"
  "math/rand"
  "fmt"
  "log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func dse_stickers_init() (s1 rand.Source, r1 rand.Rand){
  return rand.NewSource(time.Now().UnixNano()), *rand.New(s1)
}

func dse_stickers(update tgbotapi.Update, bot tgbotapi.BotAPI, s1 rand.Source, r1 rand.Rand){
  if strings.Contains(strings.ToUpper(update.Message.Text), "DSE") || strings.Contains(update.Message.Text, "文憑試"){
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