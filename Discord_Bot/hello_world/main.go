package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Details struct {
	Token string
	Url   string
}

func (d *Details) Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	d.Token = os.Getenv("DISCORD_TOKEN")
	d.Url = os.Getenv("BOT_URL")
}

func main() {
	var details Details
	details.Init()

	session, err := discordgo.New("Bot " + details.Token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}
	session.AddHandler(messageFunc)

	session.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	err = session.Open()
	if err != nil {
		log.Fatal("Error opening Discord session:", err)
	}
	defer session.Close()

	fmt.Println("The bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
func messa
geFunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Printf("%v: %v", m.Author.Username, m.Content)
	if strings.EqualFold(m.Content, "hello") {
		s.ChannelMessageSend(m.ChannelID, "World!")
	}
}
