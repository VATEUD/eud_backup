package main

import (
	backup "eud_backup/internal/app/backup"
	"flag"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
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

	flag.Parse()

	databases := flag.Args()

	backup.Start(databases)

	c := cron.New()

	c.AddFunc("@midnight", func() {
		backup.Start(databases)
	})

	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
