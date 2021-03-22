package go_content_files

import (
	"log"
	"os"
)

func CreateControllersPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/api/controllers", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeEntityControllerFile(username, projectName)
	writeServerControllerFile(projectName)
}

func writeEntityControllerFile(username, projectName string){
	var content = `package controllers

import (
	"encoding/json"
	"github.com/`+ username +`/`+ projectName +`/api/models"
	"github.com/`+ username +`/`+ projectName +`/api/repository/entity_repository"
	"github.com/`+ username +`/`+ projectName +`/api/responses"
	"io/ioutil"
	"net/http"
)

// AddEntityController is the entity function for a controller
func AddEntityController(w http.ResponseWriter, r *http.Request) {
	var entity models.Entity

	// auth := r.Header.Get("Authorization")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	entityAdded := entity_repository.AddEntity(entity)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, entityAdded)
}
`

	file, err := os.Create(projectName + "/api/controllers/entity.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeServerControllerFile(projectName string){
	var content = `package controllers

import "net/http"

// Handle the root / route to return feedback about the server to request like "Server Running..."
func ServerRunning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running..."))
}
`

	file, err := os.Create(projectName + "/api/controllers/server.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}