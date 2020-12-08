package main

import (
	backup "eud_backup/internal/app/eud_backup"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
)

// Function starts the app
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	backup.Start()

	c := cron.New()

	c.AddFunc("@every 0h1m0s", backup.Start)

	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
