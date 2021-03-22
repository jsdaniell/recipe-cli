package go_content_files

import (
	"log"
	"os"
)

func CreateCIGithubFiles(projectName string) {
	err := os.Mkdir(projectName + "/.github", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName + "/.github/workflows", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeGithubWorkflowFile(projectName)


}

func writeGithubWorkflowFile(projectName string){
	var content = `name: Heroku CI - CD

on:
  push:
    branches: [ master ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - uses: actions/setup-go@v1
        with:
          go-version: '1.14.6'
      - run: go mod vendor
      - run: git config --global user.email "putYourEmail" && git config --global user.name "${username}"
      - uses: akhileshns/heroku-deploy@v3.4.6
        with:
          heroku_api_key: \${{ secrets.HEROKU_API_KEY }}
          heroku_app_name: "${nameProject}"
          heroku_email: "putYourEmail"
`

	file, err := os.Create(projectName + "/.github/workflows/deploy.yml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}