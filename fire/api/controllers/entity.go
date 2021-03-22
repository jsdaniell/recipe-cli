package controllers

import (
	"encoding/json"
	"github.com/jsdaniell/fire/api/models"
	"github.com/jsdaniell/fire/api/repository/entity_repository"
	"github.com/jsdaniell/fire/api/responses"
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
