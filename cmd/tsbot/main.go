package main

import (
	"flag"
	"log"

	"github.com/Yol96/TSBot/internal/app/database"

	"github.com/BurntSushi/toml"

	"github.com/Yol96/TSBot/internal/app/tsbot"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/tsbot.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := tsbot.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an database connection
	err = database.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := tsbot.Start(config); err != nil {
		log.Fatal(err)
	}
}
