package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	dotenv := goDotEnvVariable("BOT_TOKEN")
	discord, _ := discordgo.New("Bot " + dotenv)

	// Register messageCreate as a callback for the messageCreate events.
	discord.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := discord.Open()
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
	discord.Close()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.Contains(m.Content, "Dolar") || strings.Contains(m.Content, "dolar") {
		// https://api.genelpara.com/embed/doviz.json
		resp, err := http.Get("https://api.genelpara.com/embed/doviz.json")
		if err != nil {
			log.Println("API'den veri çekilirken hata oluştu.")
		}
		defer resp.Body.Close()
		postBody, _ := ioutil.ReadAll(resp.Body)
		postJson := make(map[string](map[string]string))

		now := time.Now()

		json.Unmarshal(postBody, &postJson)
		log.Println(postJson["USD"]["satis"][len(postJson["USD"]["satis"])-7:])
		print_string := fmt.Sprintf("%s itibariyle 1$ = %s₺", now.Format("2006-01-02 15:04:05"), postJson["USD"]["satis"][len(postJson["USD"]["satis"])-7:])

		s.ChannelMessageSend(m.ChannelID, print_string)
	}

	if strings.Contains(m.Content, "Euro") || strings.Contains(m.Content, "euro") || strings.Contains(m.Content, "avro") || strings.Contains(m.Content, "öyrö") || strings.Contains(m.Content, "öyro") || strings.Contains(m.Content, "yöro") {
		// https://api.genelpara.com/embed/doviz.json
		resp, err := http.Get("https://api.genelpara.com/embed/doviz.json")
		if err != nil {
			log.Println("API'den veri çekilirken hata oluştu.")
		}
		defer resp.Body.Close()
		postBody, _ := ioutil.ReadAll(resp.Body)
		postJson := make(map[string](map[string]string))

		now := time.Now()

		json.Unmarshal(postBody, &postJson)
		log.Println(postJson["EUR"]["satis"])
		print_string := fmt.Sprintf("%s itibariyle 1€ = %s₺", now.Format("2006-01-02 15:04:05"), postJson["EUR"]["satis"][len(postJson["EUR"]["satis"])-7:])

		s.ChannelMessageSend(m.ChannelID, print_string)
	}

	if strings.Contains(m.Content, "Pound") || strings.Contains(m.Content, "pound") {
		// https://api.genelpara.com/embed/doviz.json
		resp, err := http.Get("https://api.genelpara.com/embed/doviz.json")
		if err != nil {
			log.Println("API'den veri çekilirken hata oluştu.")
		}
		defer resp.Body.Close()
		postBody, _ := ioutil.ReadAll(resp.Body)
		postJson := make(map[string](map[string]string))

		now := time.Now()

		json.Unmarshal(postBody, &postJson)
		log.Println(postJson["GBP"]["satis"])
		print_string := fmt.Sprintf("%s itibariyle 1£ = %s₺", now.Format("2006-01-02 15:04:05"), postJson["GBP"]["satis"][len(postJson["GBP"]["satis"])-7:])

		s.ChannelMessageSend(m.ChannelID, print_string)
	}
}
