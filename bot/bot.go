package bot

import (
	"fmt"

	"github.com/Floriansylvain/GObot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageCreate)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == config.BotPrefix+"cc" {
		s.ChannelMessageSend(m.ChannelID, "sv ?")
	}
}
