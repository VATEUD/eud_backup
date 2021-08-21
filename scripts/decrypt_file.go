package main

import (
	"eud_backup/pkg/encryption"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// main decrypts the binary file
func main() {

	// read the env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(err.Error())
		return
	}

	// check was binary provided
	if len(os.Args) < 2 {
		log.Fatalln("Please provide the file name.")
		return
	}

	fileName := os.Args[1]

	// read the file
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	// create a new zip file
	file, err := os.Create(getFileName(fileName))

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	defer file.Close()

	// construct the new cipher
	key, err := encryption.New()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	// decrypt the binary file bytes
	data, err := key.DecryptData(bytes)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	// write the zip file
	if _, err = file.Write(data); err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println("File saved!")
}

// getFileName returns the constructed zip file name
func getFileName(name string) string {
	return fmt.Sprintf("%s.zip", strings.Split(name, ".")[0])
}
