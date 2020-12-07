package commands

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/mavolin/dismock/pkg/dismock"
)

func newMock(t *testing.T) (s *discordgo.Session, m *discordgo.MessageCreate) {
	m := dismock.New(t)

	s, _ := discordgo.New("Bot abc")
	s.Client = m.Client

	return s, m
}

func TestHandlePingCommand(t *testing.T) {
	s, m = newMock(t)

}
