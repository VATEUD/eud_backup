package database

import (
	"eud_backup/pkg/config"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

// Database represents a database that's going to be backed up
type Database struct {
	Name        string
	DumpedBytes []byte
	AddedAt     time.Time
}

// commandArguments retrieves credentials from the config (.env) and later adds them to the slice which will be used to execute command
func (database *Database) commandArguments() []string {
	conf := config.RetrieveDatabaseCredentials(database.Name)
	return []string{fmt.Sprintf("-u%s", conf.Database.Username), fmt.Sprintf("-p%s", conf.Database.Password), database.Name}
}

// Dump executes the mysqldump command and stores read bytes into the databases' DumpedBytes field
func (database *Database) Dump() error {
	// construct command
	cmd := exec.Command("mysqldump", database.commandArguments()...)

	// pipe from which we'll read the output
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	// execute the command
	if err = cmd.Start(); err != nil {
		return err
	}

	// read the output
	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		return err
	}

	// store the output
	database.DumpedBytes = bytes

	return nil
}

// String returns the string representation of the database
func (database *Database) String() string {
	return database.Name
}

// FileName returns the constructed file name that will be used as a name for the sql file inside of the zip
func (database *Database) FileName() string {
	return fmt.Sprintf("%s_%s.sql", database.Name, time.Now().UTC().Format("2006_01_02"))
}
