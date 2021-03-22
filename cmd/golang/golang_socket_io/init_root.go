package golang_socket_io

import (
	"github.com/jsdaniell/recipe-cli/cmd/golang/golang_socket_io/golang_socket_content_files"
	"github.com/jsdaniell/recipe-cli/utils/go_commands"
	"log"
	"os"
)

func InitRoot(username, projectName string){
	err := os.Mkdir(projectName, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	golang_socket_content_files.CreateMainFile(username, projectName)

	go_commands.GoModInit(username, projectName)
	go_commands.GoGet("github.com/ambelovsky/gosf", projectName)

	golang_socket_content_files.CreateHTMLFile(projectName)
	golang_socket_content_files.CreateRouterPackage(username, projectName)
	golang_socket_content_files.CreateControllersPackage(projectName)
	golang_socket_content_files.CreateConfigPackage(projectName)

	go_commands.GoModTidy(projectName)
	go_commands.GoModVendor(projectName)
}