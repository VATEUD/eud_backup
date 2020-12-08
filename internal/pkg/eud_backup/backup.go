package eud_backup

import (
	"database/sql"
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
	"log"
	"time"
)

// Represents the dump directory
var dumpDir string = "dumps"

func Backup(db *sql.DB, dbname , folder string) {
	var fileName = name(dbname)

	dumper, err := mysqldump.Register(db, dumpDir + "/" + folder, fileName)

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

func name(dbname string) string {
	return fmt.Sprintf("%s-%v", dbname, time.Now().UnixNano())
}