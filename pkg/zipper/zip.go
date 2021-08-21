package zipper

import (
	"archive/zip"
	"eud_backup/pkg/database"
	"fmt"
	"os"
	"time"
)

type Archive struct {
	File *os.File
	Writer *zip.Writer
}

func (archive Archive) Zip(database *database.Database) error {
	file, err := archive.Writer.Create(fmt.Sprintf("%s.sql", database.Name))

	if err != nil {
		return err
	}

	if _, err := file.Write(database.DumpedBytes); err != nil {
		return err
	}

	return nil
}

func (archive Archive) ZipSlice(databases []*database.Database) []error {
	var errors []error

	for _, db := range databases {
		file, err := archive.Writer.Create(fmt.Sprintf("%s.sql", db.Name))

		if err != nil {
			errors = append(errors, err)
			continue
		}

		if _, err := file.Write(db.DumpedBytes); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func New() (Archive, error) {
	file, err := os.Create(fmt.Sprintf("database_backup_%s.zip", time.Now().Format("2006-01-02")))

	if err != nil {
		return Archive{}, nil
	}

	return Archive{file, zip.NewWriter(file)}, nil
}