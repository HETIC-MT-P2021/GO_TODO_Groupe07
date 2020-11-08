package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/config"
	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/models"
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
	content := strings.Split(m.Content, " ")
	command := content[0]
	ctx := context.Background()

	if command == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	if command == "!hey" {
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
	if command == "!remindme" {
		var message string

		for _, val := range content[1:] {
			message += val
		}
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		_, err := models.InsertActions(userID, message)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "DB error")

			return
		}
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")

	}

	if command == "!allremind" {
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		actions, err := models.GetUserActions(ctx, userID)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))

			return
		}
		for i := 0; i < len(actions); i++ {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", actions[i].ActionID, actions[i].Content))
		}
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}

	if command == "!lastremind" {
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		action, err := models.GetUserLastAction(userID)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))

			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", action.ActionID, action.Content))
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}

	if command == "!rmremind" {
		err := models.DeleteAction(ctx, content[1])

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))

			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Deleted Remind ID: %s", content[1]))
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
}
