package database

import (
	"eud_backup/pkg/config"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

type Database struct {
	Name        string
	DumpedBytes []byte
	AddedAt     time.Time
}

func (database *Database) commandArguments() []string {
	conf := config.RetrieveDatabaseCredentials(database.Name)
	return []string{fmt.Sprintf("-u%s", conf.Database.Username), fmt.Sprintf("-p%s", conf.Database.Password), database.Name}
}

func (database *Database) Dump() error {
	cmd := exec.Command("mysqldump", database.commandArguments()...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		return err
	}

	database.DumpedBytes = bytes

	return nil
}

func (database *Database) String() string {
	return database.Name
}

func (database *Database) FileName() string {
	return fmt.Sprintf("%s_%s.sql", database.Name, time.Now().UTC().Format("2006_01_02"))
}
