package main

import (
	"fmt"

	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/bot"
	"github.com/HETIC-MT-P2021/GO_TODO_Groupe07/config"
)

func main() {
	/*
	fmt.Println("Starting Bot")

	env, _ := godotenv.Read(".env")

	dbPort, err := strconv.ParseInt(env["DB_PORT"], 10, 64)

	if err != nil {
		panic(err)
	}

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], dbPort)

	fmt.Println("Bot is ready")
*/
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	bot.Start()

	<-make(chan struct{})
	return






}
