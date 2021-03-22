package golang_socket_content_files

import (
	"log"
	"os"
)

func CreateControllersPackage(projectName string) {
	err := os.Mkdir(projectName + "/controllers", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeEchoController(projectName)
}

func writeEchoController(projectName string){
	var content = `package controllers

import "github.com/ambelovsky/gosf"

type EchoDetail struct {
	OneThing     string ` + "`" + `json:"oneThing,omitempty"` + "`" + `
	AnotherThing struct {
		MoreDetail string ` + "`" + `json:"moreDetail,omitempty"` + "`" + `
	} ` + "`" + `json:"anotherThing,omitempty"` + "`" + `
}

type EchoRequestBody struct {
	Description string ` + "`" + `json:"description,omitempty"` + "`" + `
}

// Echo returns the passed message back to the client
func Echo(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Get request arguments and convert them to a predefined struct
	requestBody := new(EchoRequestBody)
	gosf.MapToStruct(request.Message.Body, requestBody)

	responseText := ""

	// If a detailed description was entered, send it back to the client
	if requestBody.Description != "" {
		responseText = " - " + requestBody.Description
	}

	echoDetail := &EchoDetail{
		OneThing: "this is one thing",
	}
	echoDetail.AnotherThing.MoreDetail = "and another thing..."

	return gosf.NewSuccessMessage(responseText, gosf.StructToMap(echoDetail))
}
`

	file, err := os.Create(projectName + "/controllers/echo.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

