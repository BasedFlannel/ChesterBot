package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	restTest := restGet("https://api.restful-api.dev/objects")
	log.Println(restTest)
	runBot()
}

func runBot() {

	//import a discord auth token from bot.key and instantiate a new bot with it
	authToken := loadFile("bot.key")
	sess, err := discordgo.New("Bot " + authToken)
	errorCheck(err)
	//if needed, permission int is 116736

	//add message handler, intents, and open the session
	sess.AddHandler(helloMesages)
	sess.Identify.Intents = discordgo.IntentsGuildMessages
	err = sess.Open()
	errorCheck(err)
	log.Println("The bot is listening")

	//code to hold the thread until interrupt is sent
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Primary message handler
func helloMesages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "chester") || strings.HasPrefix(m.Content, "Chester") {
		log.Println(m.Content)

		//check if there's any attatchments and loop over them
		if len(m.Attachments) > 0 {
			for _, attatchment := range m.Attachments {
				//for any image attatchemnts, log and print out the image proxy url.
				if strings.HasPrefix(attatchment.ContentType, "image") {
					log.Printf("image found: %+v\n", attatchment.ProxyURL)
					s.ChannelMessageSend(m.ChannelID, "Oh boy yummy image for me! "+attatchment.ProxyURL)
				} else {
					log.Println(attatchment.ContentType)
					s.ChannelMessageSend(m.ChannelID, "*gorps your files* Oh that's not quite as good as an image.")
				}
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Hey that's me! Gimme an image I've been so good :D")
		}
	}
}

func loadFile(filename string) string {
	data, err := os.ReadFile(filename)
	errorCheck(err)
	return (string(data))
}

// basic error check function to throw panic when nothing more is required
func errorCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}
