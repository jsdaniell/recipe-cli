package golang_cli_content_files

import (
	"log"
	"os"
)

func CreateCLIPackage(projectName string) {
	err := os.Mkdir(projectName + "/cli", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeCLISelectorFile(projectName)
	writeUserInputFile(projectName)


}

func writeCLISelectorFile(projectName string){
	var content = `package cli

import (
	"github.com/manifoldco/promptui"
	"log"
)

// SelectorCli helps to get user input showing a selector with the passed options and label.
func SelectorCli(label string, options ...string)  (string, error) {
	s := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
`

	file, err := os.Create(projectName + "/cli/selector_cli.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeUserInputFile(projectName string){
	var content = `package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
)

// An example of validate function:
//
// validate := func(input string) error {
//	 _, err := strconv.ParseFloat(input, 64)
//	 if err != nil {
// 		 return errors.New("Invalid number")
//	 }
//	 return nil
// }

// UserInput allow to get the user input with optional validate function.
func UserInput(label string, validate ...promptui.ValidateFunc) (string, error){

	var validation promptui.ValidateFunc

	if len(validate) == 0 {
		validation = promptui.ValidateFunc(func(s string) error {
			return nil
		})
	} else if len(validate) == 1 {
		validation = validate[0]
	} else if len(validate) > 1 {
		return "", errors.New("it's permited only one validation function parameter.")
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validation,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

`

	file, err := os.Create(projectName + "/cli/user_input.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}