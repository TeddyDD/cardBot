package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TeddyDD/cardBot/initiative"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	Debug bool
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.BoolVar(&Debug, "d", false, "Debug mode")
	flag.Parse()

}

func main() {
	rand.Seed(time.Now().Unix())

	if Debug {
		//d := NewDeck()
		//d.Shuffle()
		//sort.Sort(ByValue(*d))
		//for _, c := range *d {
		//fmt.Println(c.String(), "-->", c.suit, c.value)
		//}
		c := initiative.NewCombat()
		c.NewPlayer("ziutek")
		c.NewPlayer("benek")
		c.NewPlayer("waldek")
		for i := 0; i < 30; i++ {
			c.GiveCards()
			fmt.Println("Inicjatywa", i)
			c.GetInitiative()
		}
		//for _, i := range *init {
		//fmt.Println(i)
		//}
		fmt.Println("Combat deck")
		for _, i := range *c.Deck {
			fmt.Println(i)
		}
		return
	}

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

	// ping - pong
	if m.Content == "/ping" {
		if rand.Intn(50) == 42 {
			s.ChannelMessageSend(m.ChannelID, "No ile można :/")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "pong")
	} else if m.Content == "/pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}

	//if m.Content == "/card" {
	//s.ChannelMessageSend(m.ChannelID, getRandomCard())
	//}

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
