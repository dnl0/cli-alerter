package main

import (
  "log"
  "fmt"
  "flag"
  "os"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
  var message string

  flag.StringVar(&message, "m", "", "message to send")
  flag.Parse()

  if len(message) == 0 {
      fmt.Println("usage: bot.go -m \"your message\"")
      os.Exit(0)
  }

  bot, err := tgbotapi.NewBotAPI(BOT_API_TOKEN)

  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  message_tg_format := tgbotapi.NewMessage(USER_ID, message)

  bot.Send(message_tg_format)
}

