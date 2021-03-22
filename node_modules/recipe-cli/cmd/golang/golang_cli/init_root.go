package golang_cli

import (
	"github.com/jsdaniell/recipe-cli/cmd/golang/golang_cli/golang_cli_content_files"
	"github.com/jsdaniell/recipe-cli/utils/go_commands"
	"log"
	"os"
)

func InitRoot(username, projectName string){
	err := os.Mkdir(projectName, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	golang_cli_content_files.CreateMainFile(username, projectName)

	go_commands.GoModInit(username, projectName)
	go_commands.GoGet("github.com/spf13/cobra", projectName)

	golang_cli_content_files.CreateGoReleaseFile(projectName)
	golang_cli_content_files.CreatePackageJSONFile(username, projectName)
	golang_cli_content_files.CreateCIGithubFiles(projectName)
	golang_cli_content_files.CreateCmdPackage(username, projectName)
	golang_cli_content_files.CreateOtherCommandPackage(username, projectName)
	golang_cli_content_files.CreateCLIPackage(projectName)
	golang_cli_content_files.CreateUtilsPackage(projectName)

	go_commands.GoModTidy(projectName)
	go_commands.GoModVendor(projectName)
}