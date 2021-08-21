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

// Start Function reads to config file and starts the loop which backs up the database
func Start() {
	config, err := backup.ReadConfigFile()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var databases []*database.Database

	// add all databases to the databases slice
	for _, name := range config.Databases {
		databases = append(databases, database.New(name))
	}

	// start the loop
	for {
		// create a new temporary zip file
		file, err := zipper.New()

		if err != nil {
			log.Println(err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		// go through the databases slice and dump the databases
		for _, db := range databases {
			err = db.Dump()
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}

		// add the databases to the zip file
		if errs := file.Zip(databases); errs != nil {
			log.Printf("Failed to zip the files. Error:%s", err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		// close (save) the zip file
		if err = file.ZipFile.Close(); err != nil {
			log.Println(err.Error())
			time.Sleep(retryPeriod)
			continue
		}

		// upload the zip file
		if err = file.Upload(); err != nil {
			log.Println(err.Error())
			time.Sleep(retryPeriod)
			continue
		}
		
		log.Println("Backed up!")

		// sleep for 24 hours
		time.Sleep(sleepPeriod)
	}
}
