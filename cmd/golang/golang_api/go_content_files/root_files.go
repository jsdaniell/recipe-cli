package go_content_files

import (
	"log"
	"os"
)

func CreateMainFile(username, projectName string) {
	var content = `package main

import (
	"github.com/`+ username +`/` + projectName +`/api"
)

func main() {
	api.Run()
}`

	file, err := os.Create(projectName + "/main.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateEnvFile(projectName string) {
	var content = `API_PORT=9000`

	file, err := os.Create(projectName + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBuildFile(projectName string) {
	var content = `go mod vendor
go build -o bin/` + projectName + ` -v . `

	file, err := os.Create(projectName + "/build.sh")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDeployToHerokuFile(projectName string) {
	var content = `go mod vendor
git add .
git commit -a -m "Updated heroku"
git push heroku master
cd ../`

	file, err := os.Create(projectName + "/deploy_to_heroku.sh")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateHerokuYmlFile(projectName string) {
	var content = `# https://devcenter.heroku.com/articles/heroku-yml-build-manifest
# Officially unsupported, but works.
build:
  languages:
    - go

run:
  web: ` + projectName

	file, err := os.Create(projectName + "/heroku.yml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateProcfileFile(projectName string) {
	var content = `web: bin/` + projectName

	file, err := os.Create(projectName + "/Procfile")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

