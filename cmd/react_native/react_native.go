package react_native

import (
	"errors"
	"github.com/jsdaniell/recipe-cli/cli"
	"github.com/jsdaniell/recipe-cli/cmd/react_native/react_native_expo"
	"github.com/spf13/cobra"
	"log"
)

const (
	EXPO              string = "EXPO"
)

func validateProjectName(input string) error {
	if len(input) < 3 {
		return errors.New("you have to type a project name")
	}

	// TODO: Test if contains no recommended characters like - * ( ) and others.

	return nil
}


// GoCmd handles all the golang project types.
var ReactNativeCmd = &cobra.Command{
	Use:   "react-native",
	Short: "Choose the type of project:",
	Long:  `Different react-native projects options for boilerplate.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectType, err := cli.SelectorCli("Choose the type of golang project:", EXPO)
		if err != nil {
			log.Fatal(err)
		}

		projectName, err := cli.UserInput("Type the name of the project", validateProjectName)
		if err != nil {
			log.Fatal(err)
		}

		_, err = cli.UserInput("Type your username, for use as the owner of the project:", validateProjectName)
		if err != nil {
			log.Fatal(err)
		}

		if projectType == EXPO {
			react_native_expo.InitRoot(projectName)
		}
	},
}