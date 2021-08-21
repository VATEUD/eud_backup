package main

import (
	eudBackup "eud_backup/internal/app/backup"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"time"
)

type DatabasesResults map[string]string

type Data struct {
	DatabasesData DatabasesResults
	CreatedAt time.Time
}

// Function starts the app
func main() {

	eudBackup.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
