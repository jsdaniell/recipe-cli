package golang_cli_content_files

import (
	"log"
	"os"
)

func CreateOtherCommandPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/cmd/other_command", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeOtherCommandFile(username, projectName)
}

func writeOtherCommandFile(username, projectName string){
	var content = `package other_commands

import (
	"errors"
	"github.com/` + username + `/` + projectName + `/cli"
	"github.com/spf13/cobra"
)

 func validateProjectName(input string) error {
		 if len(input) < 3 {
		 	return errors.New("you have to type a project name")
		 }

		 // TODO: Test if contains no recommended characters like - * ( ) and others.

		 return nil
 }


// OtherCommand handles other command, customize it!.
var OtherCommand = &cobra.Command{
	Use:   "othercommand",
	Short: "Other command short description",
	Long:  ` + "`" + `Other command long description` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		// run your different command
	},
}
`

	file, err := os.Create(projectName + "/cmd/other_command/other_command.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}