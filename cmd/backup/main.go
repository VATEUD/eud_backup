package main

import (
	"eud_backup/internal/app/backup"
	"github.com/joho/godotenv"
	"log"
)

// Function starts the app
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	eudbackup.Start()
}
