package reminds

import (
	"context"
	"github.com/bwmarrin/discordgo"

)

type Reminds interface {
    HandleGetRemindsCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx context.Context)
	HandleGetLastRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate)
	HandleDeleteRemindCommand(s *discordgo.Session, m *discordgo.MessageCreate, content []string, ctx context.Context)
	HandleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate)
	ChannelMessageSend(string, string) (*discordgo.Message, error)
	HandleDefaultCommand(s *discordgo.Session, m *discordgo.MessageCreate)
}