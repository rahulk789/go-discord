package main

import (
  "fmt"
  "go-discord/bot"
  "go-discord/config"
)

func main() {
  err := config.ReadConfig()

  if err != nil {
  	fmt.Println(err.Error())
  	return
  }

  bot.Start()

  <-make(chan struct{})
  return
}
