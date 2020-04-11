package main

import (
	"log"
  //"os"
	"github.com/go-ego/gse"
  "strings"
  "github.com/go-telegram-bot-api/telegram-bot-api"
)
var (
  seg gse.Segmenter
)

const replyText = "%s完又%s\n日%s夜%s\n一%s再%s\n死%s爛%s"

/*func nlp(){
  seg.LoadDict("dictionary/dict_big.txt")
  bot := authorization()
  updates := configure_webhook(bot)
  
  //s1, r1 := dse_stickers_init()
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for update := range updates {
    if update.Message == nil { // ignore any non-Message Updates
			continue
	  }
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    temp := []byte(update.Message.Text)

    a := seg.ModeSegment(temp)
    log.Print(seg.String(temp))
    log.Print(len(a))
  
    for i := 0; i < len(a); i++ {
      log.Print(a[i].Token().Pos())
      log.Print(a[i].Token().Text())
      //log.Print(len(a[i].Token().Text()))
      if (strings.Contains(a[i].Token().Pos(), "v") || strings.Contains(a[i].Token().Pos(), "zg")) && len(a[i].Token().Text()) == 3 {
        log.Print("Contains verb")
        message := strings.ReplaceAll(replyText, "%s", a[i].Token().Text())
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			  msg.ReplyToMessageID = update.Message.MessageID
			  bot.Send(msg)
        log.Print("Sent!")
        break
      }
    }
  }
}*/

func nlp_message(bot tgbotapi.BotAPI, update tgbotapi.Update){
    temp := []byte(update.Message.Text)

    a := seg.ModeSegment(temp)
    log.Print(seg.String(temp))
    log.Print(len(a))
  
    for i := 0; i < len(a); i++ {
      log.Print(a[i].Token().Pos())
      log.Print(a[i].Token().Text())
      //log.Print(len(a[i].Token().Text()))
      if (strings.Contains(a[i].Token().Pos(), "v") || strings.Contains(a[i].Token().Pos(), "zg")) && len(a[i].Token().Text()) == 3 {
        log.Print("Contains verb")
        message := strings.ReplaceAll(replyText, "%s", a[i].Token().Text())
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			  msg.ReplyToMessageID = update.Message.MessageID
			  bot.Send(msg)
        log.Print("Sent!")
        break
      }
    }
}
