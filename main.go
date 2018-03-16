package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	Cards []string
)

func CreateCards() {
	color := []string{"♥", "♦", "♠", "♣"}
	value := []string{"9", "10", "J", "Q", "K", "A"}
	for _, c := range color {
		for _, v := range value {
			Cards = append(Cards, fmt.Sprintf("%s %s", c, v))
		}
	}
	Cards = append(Cards, "Czarny Joker")
	Cards = append(Cards, "Czerwony Joker")
}

func getRandomCard() string {
	return Cards[rand.Intn(len(Cards))]
}

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().Unix())

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	CreateCards()
	CreateWeather()
	fmt.Println("creted card deck")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot == true {
		return
	}

	if m.Content == "/card" {
		s.ChannelMessageSend(m.ChannelID, getRandomCard())
	}

	if m.Content == "/zima" {
		s.ChannelMessageSend(m.ChannelID, GetWeatherReport("zima"))
	}
	if m.Content == "/wiosna" {
		s.ChannelMessageSend(m.ChannelID, GetWeatherReport("wiosna"))
	}
	if m.Content == "/jesień" {
		s.ChannelMessageSend(m.ChannelID, GetWeatherReport("jesień"))
	}
	if m.Content == "/lato" {
		s.ChannelMessageSend(m.ChannelID, GetWeatherReport("lato"))
	}

	if m.Content == "/polowanie" {
		s.ChannelMessageSend(m.ChannelID, GetAnimal())
	}
}
