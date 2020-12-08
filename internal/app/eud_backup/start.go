package eud_backup

import (
	"eud_backup/internal/pkg/eud_backup/sql"
	"fmt"
	"log"
	backup "eud_backup/internal/pkg/eud_backup"
	"os"
	"time"
)

func Start() {

	databases := []string{
		"eud_site", "eud_central", "myvatsim", "central",
	}

	now := time.Now().UTC()

	folder := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())

	for _, name := range databases {
		db, err := sql.Connect(name)

		if err != nil {
			log.Println(err)
			return
		}

		if _, err := os.Stat("dumps/" + folder); os.IsNotExist(err) {

			err := createDirctory(fmt.Sprintf("dumps/%s", folder), 0775)

			if err != nil {
				log.Println(err)

				continue
			}

		}

		backup.Backup(db, name, folder)
	}

}

func createDirctory(name string, perms os.FileMode) error {
	err := os.Mkdir(name, perms)

	if err != nil {
		return err
	}

	return nil
}