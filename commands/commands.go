package commands

import (
	"context"
	"fmt"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/models"
	"github.com/bwmarrin/discordgo"
)

// HandlePingCommand is the basic command to know if the server is Up
func HandlePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
}

// HandleAddRemindCommand insert a new remind for a user
func HandleAddRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string) {
	var message string

	for _, val := range content[1:] {
		message += fmt.Sprintf("%s ", val)
	}
	userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

	_, err := models.InsertRemind(userID, message)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "DB error")
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		return
	}
	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}

// HandleGetRemindsCommand get all the user's reminds
func HandleGetRemindsCommand(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate) {
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

// HandleGetLastRemindCommand get the user's last remind
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

// HandleDeleteRemindCommand handles deletion of user's remind
func HandleDeleteRemindCommand(ctx context.Context, s *discordgo.Session, m *discordgo.MessageCreate, content []string) {
	userID := fmt.Sprintf("%s-%s", m.Author.Username, m.Author.ID)

	err := models.DeleteRemind(ctx, content[1], userID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err))
		s.MessageReactionAdd(m.ChannelID, m.ID, "❌")
		return
	}
	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}

// HandleHelpCommand handles the help message
func HandleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpContent := `Hello there ! 
	 Here are the commands you can enter:
	 - !remindme 'string', will save a reminder for you with the content after
	 - !allremind, will gets all your reminder
	 - !lastremind, will get the last remind you entered
	 - !rmremind 'remind_ID', will delete selected remind`

	s.ChannelMessageSend(m.ChannelID, helpContent)
}

// HandleDefaultCommand handles commands that couldn't be recognized
func HandleDefaultCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Couldn't understand your request, type !todohelp for more informations")
}
