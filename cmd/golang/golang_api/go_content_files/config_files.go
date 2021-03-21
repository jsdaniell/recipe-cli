package go_content_files

import (
	"log"
	"os"
)

func CreateConfigPackage(projectName string) {
	err := os.Mkdir(projectName + "/config", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}


	writeConfigFile(projectName)
}

func writeConfigFile(projectName string){
	var content = `package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 9000
)

func Load(){

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}
}
`

	file, err := os.Create(projectName + "/config/config.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

