package eud_backup

import (
	"eud_backup/internal/pkg/eud_backup/sql"
	"log"
)

func Start() {

	databases := []string{
		"eud_site", "eud_central",
	}

	for _, name := range databases {
		db, err := sql.Connect(name)

		if err != nil {
			log.Println(err)
			return
		}

		db.Close()
	}

}
