package main

import (
  "log"
  //"strconv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func dictator(bot tgbotapi.BotAPI, update tgbotapi.Update, chatid int64) {
  
		if update.Message.From.UserName == "Art_3mis"{
      msg := tgbotapi.NewDeleteMessage(chatid, update.Message.MessageID)
      _, err := bot.DeleteMessage(msg)
      log.Print(err)
      //log.Printf("Deleted yay!")
	  }
}