package go_content_files

import (
	"log"
	"os"
)

func CreateApiPackage(username, projectName string) {


	err := os.Mkdir(projectName + "/api", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeApiServerFile(username, projectName)


}

func writeApiServerFile(username, projectName string){
	var content = `package api

import (
	"fmt"
	"github.com/`+ username +`/`+ projectName +`/api/router"
	"github.com/`+ username +`/`+ projectName +`/config"
	"log"
	"net/http"
)

// Run initialize the server.
func Run() {
	config.Load()
	fmt.Printf("Listening... localhost:%d", config.PORT)
	listen(config.PORT)
}

// Configure listen port.
func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
`

	file, err := os.Create(projectName + "/api/server.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}