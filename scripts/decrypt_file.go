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

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println(err.Error())
		return
	}

	if len(os.Args) < 2 {
		log.Fatalln("Please provide the file name.")
		return
	}

	fileName := os.Args[1]

	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	file, err := os.Create(getFileName(fileName))

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	defer file.Close()

	key, err := encryption.New()

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	data, err := key.DecryptData(bytes)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	if _, err = file.Write(data); err != nil {
		log.Fatalln(err.Error())
		return
	}

	log.Println("File saved!")
}

func getFileName(name string) string {
	return fmt.Sprintf("%s.zip", strings.Split(name, ".")[0])
}
