package main

import (
	"eud_backup/internal/app/backup"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

// Function starts the app
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	eudbackup.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
