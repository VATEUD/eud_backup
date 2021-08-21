package eudBackup

import (
	"eud_backup/internal/pkg/backup"
	"eud_backup/pkg/database"
	"eud_backup/pkg/zipper"
	"log"
)

func Start() {
	config, err := backup.ReadConfigFile()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var databases []*database.Database

	for _, name := range config.Databases {
		databases = append(databases, database.New(name))
	}

	for _, db := range databases {
		if err := db.Dump(); err != nil {
			return
		}
	}

	file, err := zipper.New()

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.File.Close()
	defer file.Writer.Close()

	for _, db := range databases {
		err := file.Zip(db)

		if err != nil {
			log.Fatalln(err.Error())
			return
		}
	}
}
