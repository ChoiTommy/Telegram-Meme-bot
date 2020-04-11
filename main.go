package main

import (
  "math/rand"
  "time"
  "log"
)

func main() {
  //init
  seg.LoadDict("dictionary/dict_big.txt")
  bot := authorization()
  updates := configure_webhook(bot)

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for update := range updates {
    if update.Message == nil { // ignore any non-Message Updates
			continue
	  }
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    nlp_message(bot, update)
    dse_stickers(update, bot, s1, *r1)
  }
}