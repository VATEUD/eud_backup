package examples

import (
	"eud_backup/internal/pkg/backup"
	"eud_backup/pkg/database"
	"fmt"
	"log"
	"os"
)

func main() {
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

		file, err := os.CreateTemp("backups", fmt.Sprintf("%s.sql", db.Name))

		if err != nil {
			return
		}

		_, err = file.Write(db.DumpedBytes)

		if err != nil {
			return
		}

	}
}
