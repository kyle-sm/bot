package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var token string
var buffer = make([][]byte, 0)

func init() {
	flag.StringVar(&token, "token", "", "Token to start this bot")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	} else {
		fmt.Println("Miku is online.")
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			//fmt.Println() // Println will add back the final '\n'
			_, _ = dg.ChannelMessageSend("435224851423428622", scanner.Text())
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

	return
}

//[435224851423428620|435224851423428622]
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(strings.ToLower(m.Message.Content), "miku play despacito") {
		_, _ = s.ChannelMessageSend(m.Message.ChannelID, "https://www.youtube.com/watch?v=40qJapBsOp4")
	}

	fmt.Printf("%s: %s \t[%s|%s]\n", m.Message.Author.Username, m.Message.Content, m.Message.GuildID, m.Message.ChannelID)

	args := strings.Split(m.Message.Content, " -")

	switch args[0] {
	case "!miku":
		_, _ = s.ChannelMessageSend(m.Message.ChannelID, "im miku")
	case "!get":
		if len(args) < 2 {
			s.ChannelMessageSend(m.Message.ChannelID, VocaRand())
		} else {
			s.ChannelMessageSend(m.Message.ChannelID, VocaGet(args[1]))
		}
	case "!spin":
		s.ChannelMessageSend(m.Message.ChannelID, Roulette())
	case "!snowflakerole":
		if len(args) != 3 {
			s.ChannelMessageSend(m.Message.ChannelID, "That's not how that command works!")
		} else {
			SnowflakeRole(s, m.Message.GuildID, m.Message.Author.ID, args[1], args[2])
		}
	}
}
