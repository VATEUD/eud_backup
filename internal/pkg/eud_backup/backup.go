package eud_backup

import (
	"database/sql"
	"github.com/JamesStewy/go-mysqldump"
	"log"
)

// Represents the dump directory
var dumpDir string = "dumps"

func Backup(db *sql.DB, dbname, folder string) {

	dumper, err := mysqldump.Register(db, dumpDir+"/"+folder, dbname)

	if err != nil {
		log.Println(err)
		return
	}

	result, err := dumper.Dump()

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("File saved (%s)", result)

	dumper.Close()
}
