package golang_cli_content_files

import (
	"log"
	"os"
)

func CreateMainFile(username, projectName string) {
	var content = `package main

import (
	"github.com/`+ username +`/`+ projectName +`/cmd"
)

func main() {
	cmd.Execute()
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

func CreateGoReleaseFile(projectName string) {
	var content = `# This is an example .goreleaser.yml file with some sane defaults.
# Make sure to check the documentation at http://goreleaser.com
before:
  hooks:
    # You may remove this if you don't use go modules.
    - go mod tidy
    # you may remove this if you don't need go generate
    - go generate ./...
project_name: `+ projectName +`
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
checksum:
  name_template: 'checksums.txt'
snapshot:
  name_template: "{{ .Tag }}-next"
changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
`

	file, err := os.Create(projectName + "/.goreleaser.yml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePackageJSONFile(username, projectName string) {
	var content = `{
  "name": "` + projectName + `",
  "version": "1.0.0",
  "description": "Description of your project",
  "main": "index.js",
  "scripts": {
    "postinstall": "go-npm install",
    "preuninstall": "go-npm uninstall"
  },
  "goBinary": {
    "name": "` + projectName + `",
    "path": "./bin",
    "url": "https://github.com/` + username + `/` + projectName + `/releases/download/v{{version}}/recipe-cli_{{version}}_{{platform}}_{{arch}}.tar.gz"
  },
  "author": "` + username + `",
  "license": "ISC",
  "bugs": {
    "url": "https://github.com/` + username + `/` + projectName + `/issues"
  },
  "homepage": "https://github.com/` + username + `/` + projectName + `#readme",
  "dependencies": {
    "go-npm": "^0.1.9",
    "recipe-cli": "latest"
  }
}

`

	file, err := os.Create(projectName + "/package.json")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}
