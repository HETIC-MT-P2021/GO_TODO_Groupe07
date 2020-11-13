package bot

import (
	"testing"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/config"
	"github.com/bwmarrin/discordgo"
)

func generateFakeMessage(content string, authorID string) discordgo.MessageCreate {
	var messageCreate discordgo.MessageCreate
	var message discordgo.Message
	var author discordgo.User

	message.Content = content
	author.ID = authorID
	message.Author = &author
	messageCreate.Message = &message

	return messageCreate
}

func TestShouldSkipMessage(t *testing.T) {
	BotID = "420"
	config.BotPrefix = "!"

	messageWithoutCommand := generateFakeMessage("Command", "666")
	messageWithBadSuffix := generateFakeMessage("%Command", "669")
	messageFromBotWithoutCommand := generateFakeMessage("Command", BotID)
	messageFromBot := generateFakeMessage("!Command", BotID)
	messageWithCommand := generateFakeMessage("!Command", "911")

	tables := []struct {
		message discordgo.MessageCreate
		result  bool
	}{
		{messageWithoutCommand, true},
		{messageWithBadSuffix, true},
		{messageFromBotWithoutCommand, true},
		{messageFromBot, true},
		{messageWithCommand, false},
	}

	for _, table := range tables {
		result := ShouldSkipMessage(&table.message)

		if result != table.result {
			t.Errorf("Test of ShouldBotSkipMessage for %s, %s was incorrect, got: %t, want: %t.", table.message.Content, table.message.Author.ID, result, table.result)
		}
	}
}
