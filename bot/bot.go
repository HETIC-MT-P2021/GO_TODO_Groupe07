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
			message += fmt.Sprintf("%s ", val)
		}
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		_, err := models.InsertReminds(userID, message)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "DB error")
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")

	}

	if command == "!allremind" {
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		reminds, err := models.GetUserReminds(ctx, userID)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}
		for i := 0; i < len(reminds); i++ {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", reminds[i].RemindID, reminds[i].Content))
		}
	}

	if command == "!lastremind" {
		userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

		remind, err := models.GetUserLastRemind(userID)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", remind.RemindID, remind.Content))
	}

	if command == "!rmremind" {
		err := models.DeleteRemind(ctx, content[1])

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
			s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
			return
		}
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
}
