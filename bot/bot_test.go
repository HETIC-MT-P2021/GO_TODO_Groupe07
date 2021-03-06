package bot

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/config"
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

// TestStart tests the Start() function without any arguments.  This should return
// a valid Session{} struct and no errors.
func TestStart(t *testing.T) {

	_, err := discordgo.New()
	if err != nil {
		t.Errorf("New() returned error: %+v", err)
	}
}


// TestInvalidToken tests the New() function with an invalid token
func TestInvalidToken(t *testing.T) {
	d, err := discordgo.New("asjkldhflkjasdh")
	if err != nil {
		t.Fatalf("New(InvalidToken) returned error: %+v", err)
	}

	// New with just a token does not do any communication, so attempt an api call.
	_, err = d.UserSettings()
	if err == nil {
		t.Errorf("New(InvalidToken), d.UserSettings returned nil error.")
	}
}

// TestNewToken tests the New() function with a Token.
func TestNewToken(t *testing.T) {

	if config.Token == "" {
		t.Skip("Skipping New(token), token not set")
	}

	d, err := discordgo.New(config.Token)
	if err != nil {
		t.Fatalf("New(envToken) returned error: %+v", err)
	}

	if d == nil {
		t.Fatal("New(envToken), d is nil, should be Session{}")
	}

	if d.Token == "" {
		t.Fatal("discordgo.New(config.Token), d.Token is empty, should be a valid Token.")
	}
}


func TestGetParamsFromMessage(t *testing.T) {
	BotID = "420"
	config.BotPrefix = "!"
	ctxTest := context.Background()

	messageWithoutCommand := generateFakeMessage("Command yes", "666")

	content, command, ctx := getParamsFromMessage(&messageWithoutCommand)

	if content[1] != "yes" || command != "Command" || ctx != ctxTest {
		t.Errorf("Test of getParamsFromMessage for %s was incorrect.", content)
	}
}

