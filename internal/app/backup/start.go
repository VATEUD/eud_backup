package eudbackup

import (
	"eud_backup/internal/pkg/backup"
	"eud_backup/pkg/web"
	"log"
)

// Start Function reads to config file and starts the loop which backs up the database
func Start() {
	// start the loop
	go backup.Start()

	server := web.New()
	if err := server.Start(); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
