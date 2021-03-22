package golang_cli_content_files

import (
	"log"
	"os"
)

func CreateUtilsPackage(projectName string) {
	err := os.Mkdir(projectName + "/utils", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName + "/utils/go_commands", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName + "/utils/shell_commands", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeGoMod(projectName)
	writeShellCommands(projectName)
}

func writeGoMod(projectName string){
	var content = `package go_commands

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
)

func GoModInit(username, projectName string){
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "init", "github.com/" + username + "/" + projectName},
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoModTidy(projectName string){
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "tidy"},
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoGet(packageName, projectName string){
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"get", "-u", packageName},
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func GoModVendor(projectName string){
	cmd := execute.ExecTask{
		Command:     "go",
		Args:        []string{"mod", "vendor"},
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}
`

	file, err := os.Create(projectName + "/utils/go_commands/go_mod.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeShellCommands(projectName string){
	var content = `package shell_commands

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
)

func ExecuteSh(file, projectName string){
	cmd := execute.ExecTask{
		Command:     "sh",
		Args:        []string{file},
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}

func ExecuteShellCommand(command string, projectName string, args ...string){
	cmd := execute.ExecTask{
		Command:     command,
		Args:        args,
		StreamStdio: false,
		Cwd: projectName,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %s", res.Stderr)
}
`

	file, err := os.Create(projectName + "/utils/go_commands/shell_commands.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}