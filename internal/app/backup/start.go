package eudBackup

import (
	"eud_backup/internal/pkg/backup"
	"eud_backup/pkg/database"
	"eud_backup/pkg/zipper"
	"log"
	"time"
)

const (
	retryPeriod = time.Minute * 5
	sleepPeriod = time.Hour * 24
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

	for {
		file, err := zipper.New()

		if err != nil {
			log.Println(err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		for _, db := range databases {
			err = db.Dump()
			if err != nil {
				log.Println(err.Error())
				time.Sleep(retryPeriod)
				continue
			}
		}

		if errs := file.Zip(databases); errs != nil {
			log.Printf("Failed to zip the files. Error:%s", err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		if err = file.ZipFile.Close(); err != nil {
			log.Println(err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		if err := file.Upload(); err != nil {
			log.Println(err.Error())
		}

		time.Sleep(sleepPeriod)
	}
}
