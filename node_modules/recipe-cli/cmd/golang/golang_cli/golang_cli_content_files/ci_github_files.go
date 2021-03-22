package golang_cli_content_files

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
	var content = `name: goreleaser

on:
  push:
    tags:
      - '*'

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      -
        name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      -
        name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16
      -
        name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.TOKEN_GO_RELEASE }}
`

	file, err := os.Create(projectName + "/.github/workflows/release.yml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}