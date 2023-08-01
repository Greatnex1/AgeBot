package main

import (
	"AgeBot/bot"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	bot.AgeLoad()
}
