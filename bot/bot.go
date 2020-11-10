package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/commands"
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

func shouldSkipMessage(m *discordgo.MessageCreate) bool {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return true
		}
		return false
	}
	return true
}

func getParamsFromMessage(m *discordgo.MessageCreate) ([]string, string, context.Context) {
	content := strings.Split(m.Content, " ")
	command := content[0]
	ctx := context.Background()

	return content, command, ctx
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if shouldSkipMessage(m) {
		return
	}

	content, command, ctx := getParamsFromMessage(m)

	switch command {
	case "!ping":
		commands.HandlePingCommand(s, m)

	case "!remindme":
		commands.HandleAddRemindCommand(s, m, content)

	case "!allremind":
		commands.HandleGetRemindsCommand(s, m, ctx)

	case "!lastremind":
		commands.HandleGetLastRemindCommand(s, m)

	case "!rmremind":
		commands.HandleDeleteRemindCommand(s, m, content, ctx)

	case "!help":
		commands.HandleHelpCommand(s, m)

	default:
		commands.HandleDefaultCommand(s, m)
	}
}
