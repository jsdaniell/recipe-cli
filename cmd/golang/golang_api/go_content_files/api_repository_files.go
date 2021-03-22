package go_content_files

import (
	"log"
	"os"
)

func CreateRepositoryPackage(username, projectName, projectDatabase string) {
	err := os.Mkdir(projectName + "/api/repository", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName + "/api/repository/entity_repository", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeEntityRepositoryFile(username, projectName, projectDatabase)
}

func writeEntityRepositoryFile(username,projectName, projectDatabase string) {
	var contentFirebase = `package entity_repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/`+ username +`/`+ projectName +`/api/db"
	"github.com/`+ username +`/`+ projectName +`/api/models"
	"github.com/`+ username +`/`+ projectName +`/api/utils/json_utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddEntity(name string, description string) (*firestore.WriteResult, error) {
	client := db.FirestoreClient()
	defer client.Close()

	var entityModel models.Entity

	entityModel.Name = name
	entityModel.Description = description

	entityCollection := client.Collection("entities")

	// TODO: Change Keys of Suite Model to lowerCase
	

	lowerCaseJson, err := json_utility.StructToLowerCaseJson(suiteModel)
	if err != nil {
		return nil, err
	}


	doc, err := entityCollection.Doc(name).Get(context.Background())
	if status.Code(err) == codes.NotFound {
		res, err := entityCollection.Doc(name).Set(context.Background(), lowerCaseJson)
		if err != nil {
			fmt.Errorf("error on registre new suite")
		}

		return res, nil
	} else {
		if doc.Exists() {
			return nil, fmt.Errorf(` + "`" + `the %q suite already exists` + "`" + `, name)
		} else {
			return nil, fmt.Errorf("error on create suite")
		}
	}
}
`
	var contentMongoDB = `package entity_repository

import (
	"context"
	"errors"
	"github.com/`+ username +`/`+ projectName +`/api/db"
	"github.com/`+ username +`/`+ projectName +`/api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func AddEntity(entity models.Entity) error {

	var exists models.Entity

	client, err := db.ConnectDB()
	if err != nil {
		return err
	}

	filter := bson.M{"name": entity.Name}

	errs := client.Collection(db.EntitiesCollection).FindOne(context.Background(), filter).Decode(&exists)
	if errs == nil {
		return errs
	}

	if (exists != models.Entity{}) {
		return errors.New("this entity already exists")
	}

	_, err = client.Collection(db.EntitiesCollection).InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEntity(name string) error {
	client, err := db.ConnectDB()
	if err != nil {
		return err
	}

	filter := bson.M{"name": "entityName"}

	errs := client.Collection(db.EntitiesCollection).FindOneAndDelete(context.Background(), filter)
	if errs.Err() == nil {
		return errors.New("do not exists")
	}

	return nil
}
`

	var contentNoSelection = `package entity_repository

import (
	"github.com/`+ username +`/`+ projectName +`/api/models"
)

// Put your database functions of specific entity here.
func AddEntity(entity models.Entity) models.Entity {
return models.Entity{}
}
`

	file, err := os.Create(projectName + "/api/repository/entity_repository/entity_repository.go")
	if err != nil {
		log.Fatal(err)
	}

	if projectDatabase == "NoSelection" {
		_, err = file.WriteString(contentNoSelection)
		if err != nil {
			log.Fatal(err)
		}
	}

	if projectDatabase == "Firebase" {
		_, err = file.WriteString(contentFirebase)
		if err != nil {
			log.Fatal(err)
		}
	}

	if projectDatabase ==  "MongoDB" {
		_, err = file.WriteString(contentMongoDB)
		if err != nil {
			log.Fatal(err)
		}
	}
}