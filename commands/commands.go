package commands

import (
	"context"
	"fmt"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/models"
	"github.com/bwmarrin/discordgo"
)

func HandlePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
}

func HandleAddRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string) {
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

func HandleGetRemindsCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx context.Context) {
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

func HandleGetLastRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

	remind, err := models.GetUserLastRemind(userID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", remind.RemindID, remind.Content))
}

func HandleDeleteRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, ctx context.Context) {
	userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

	remind, err := models.GetUserLastRemind(userID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Remind ID: %d, Content: %s", remind.RemindID, remind.Content))
}

func HandleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpContent := "Hello there ! \n Here are the commands you can enter:\n - !remindme 'string', will save a reminder for you with the content after\n - !allremind, will gets all your reminder\n - !lastremind, will get the last remind you entered\n - !rmremind 'remind_ID', will delete selected remind"

	s.ChannelMessageSend(m.ChannelID, helpContent)
}

func HandleDefaultCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Couldn't understand your request, type !todohelp for more information")
}
