package testdiscord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NzY2NjYzNzE2ODU4MzYzOTY0.X4mpOg.yoCvFqqyTzKMYxgq3MrvZG4ugoo"
var BotID string
func main()  {
	dg, err := discordgo.New("Bot" + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())

	}
	BotID = u.ID

	//dg.AddHandler(messageHandler)
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return

	}

	fmt.Println("Bot is running")

	<- make(chan struct {})
	return
}

/*func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {

	if.m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend("766638109689118723", "pong")

	}
	fmt.Println(m.Content)
	
}

 */
