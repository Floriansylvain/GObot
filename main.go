package main

import (
	"fmt"

	"github.com/Floriansylvain/GObot/bot"
	"github.com/Floriansylvain/GObot/config"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	select {}
}
