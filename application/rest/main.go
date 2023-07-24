package main

import (
	"bistleague-be/model/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	app, err := applicationDelegate(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err = startServer(cfg, app); err != nil {
		log.Fatal(err)
	}
}
