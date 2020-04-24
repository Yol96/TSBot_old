package tsbot

import (
	"log"
	"time"

	"github.com/darfk/ts3"
)

type TeamspeakBot struct {
	config *Config
}

func New(config *Config) *TeamspeakBot {
	return &TeamspeakBot{
		config: config,
	}
}

func Start(config *Config) error {
	// Make a new ts3 client
	client, err := ts3.NewClient(config.ServerAddress + config.ServerPort)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Query login
	_, err = client.Exec(ts3.Login(config.QueryLogin, config.QueryPassword))
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Use first Virtual Server
	_, err = client.Exec(ts3.Use(1))
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Bot is starting with next parameters: ", config)

	AddClientListeners(client)

	for {
		time.Sleep(500 * time.Millisecond)
	}
}
