package golang_api

import (
	"github.com/jsdaniell/recipe-cli/cmd/golang/golang_api/go_content_files"
	"github.com/jsdaniell/recipe-cli/utils/go_commands"
	"github.com/jsdaniell/recipe-cli/utils/shell_commands"
	"log"
	"os"
)

func InitRoot(username, projectName, projectDatabase string) {
	err := os.Mkdir(projectName, os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	go_content_files.CreateMainFile(username, projectName)

	go_commands.GoModInit(username, projectName)

	go_content_files.CreateApiPackage(username, projectName)
	go_content_files.CreateResponsesPackage(projectName)
	go_content_files.CreateModelsPackage(projectName, projectDatabase)
	go_content_files.CreateDbPackage(projectName, projectDatabase)
	go_content_files.CreateRepositoryPackage(username,projectName, projectDatabase)
	go_content_files.CreateControllersPackage(username, projectName)
	go_content_files.CreateMiddlewaresPackage(username, projectName)
	go_content_files.CreateRouterPackage(username, projectName)
	go_content_files.CreateUtilsPackage(username, projectName)
	go_content_files.CreateConfigPackage(projectName)

	go_content_files.CreateBuildFile(projectName)
	go_content_files.CreateCIGithubFiles(projectName)
	go_content_files.CreateDeployToHerokuFile(projectName)
	go_content_files.CreateEnvFile(projectName)
	go_content_files.CreateHerokuYmlFile(projectName)
	go_content_files.CreateProcfileFile(projectName)

	go_commands.GoModTidy(projectName)
	shell_commands.ExecuteSh("build.sh", projectName)
}
