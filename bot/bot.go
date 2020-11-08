package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Tout roule!!")
}


func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}

	}

	if m.Content == "pongg" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	if m.Content == "hey" {
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
}
