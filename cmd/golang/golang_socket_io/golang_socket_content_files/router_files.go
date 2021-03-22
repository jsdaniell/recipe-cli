package golang_socket_content_files

import (
	"log"
	"os"
)

func CreateRouterPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/router", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeRouterFile(username, projectName)
}

func writeRouterFile(username, projectName string){
	var content = `package router

import (
	"github.com/ambelovsky/gosf"
	"github.com/`+ username +`/`+ projectName +`/controllers"
)

// RegisterRoutes maps controller functions to endpoints
func RegisterRoutes() {
	gosf.Listen("echo", controllers.Echo)
}
`

	file, err := os.Create(projectName + "/router/router.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

