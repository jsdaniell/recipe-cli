package golang

import (
	"errors"
	"github.com/jsdaniell/recipe-cli/cli"
	"github.com/jsdaniell/recipe-cli/cmd/golang/golang_api"
	"github.com/jsdaniell/recipe-cli/cmd/golang/golang_cli"
	"github.com/spf13/cobra"
	"log"
)

const (
	API              string = "API"
	CLI              string = "CLI"
	MongoDBDatabase         = "MongoDB"
	FirebaseDatabase        = "Firebase"
	NoSelection             = "NoSelection"
)

 func validateProjectName(input string) error {
		 if len(input) < 3 {
		 	return errors.New("you have to type a project name")
		 }

		 // TODO: Test if contains no recommended characters like - * ( ) and others.

		 return nil
 }


// GoCmd handles all the golang project types.
var GoCmd = &cobra.Command{
	Use:   "golang",
	Short: "Choose the type of project:",
	Long:  `Different golang projects options for boilerplate.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectType, err := cli.SelectorCli("Choose the type of golang project:", API, CLI)
		if err != nil {
			log.Fatal(err)
		}

		projectName, err := cli.UserInput("Type the name of the project", validateProjectName)
		if err != nil {
			log.Fatal(err)
		}

		username, err := cli.UserInput("Type your username, naturally the package of go api will be (github.com/username/project_name):", validateProjectName)
		if err != nil {
			log.Fatal(err)
		}

		if projectType == API {
			projectDatabase, err := cli.SelectorCli("Choose the database used on project:", MongoDBDatabase, FirebaseDatabase, NoSelection)
			if err != nil {
				log.Fatal(err)
			}

			golang_api.InitRoot(username, projectName, projectDatabase)
		}

		if projectType == CLI {
			golang_cli.InitRoot(username, projectName)
		}
	},
}