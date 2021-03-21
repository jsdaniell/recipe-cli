package go_content_files

import (
	"github.com/jsdaniell/recipe-cli/cmd/golang"
	"log"
	"os"
)

func CreateModelsPackage(projectName, projectDatabase string) {


	err := os.Mkdir(projectName + "/api/models", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeEntityModelFile(projectName, projectDatabase)

	writeTokenModelFile(projectName, projectDatabase)


}

func writeEntityModelFile(projectName, projectDatabase string){
	var contentFirebase = `// Entity model used to receive request and send JSON response, is used on authentication controller.
type Entity struct {
	ID            string   ` + "`" + `json:"_id"` + "`" + `
	Name          string   ` + "`" + `json:"name, omitempty"` + "`" + `
	Description   string   ` + "`" + `json: "description"` + "`" + `
}
`
	var contentMongoDB = `package models

// Entity model used to receive request and send JSON response, is used on authentication controller.
type Entity struct {
	ID          string      ` + "`" + `json:"_id,omitempty" bson:"_id,omitempty"` + "`" + `
	Name        string      ` + "`" + `json:"name,omitempty" bson:"name,omitempty"` + "`" + `
	Description string      ` + "`" + `json:"description" bson:"description,omitempty"` + "`" + `
}
`

	file, err := os.Create(projectName + "/api/models/Entity.go")
	if err != nil {
		log.Fatal(err)
	}

	if projectDatabase == golang.NoSelection || projectDatabase == golang.FirebaseDatabase {
		_, err = file.WriteString(contentFirebase)
		if err != nil {
			log.Fatal(err)
		}
	}

	if projectDatabase ==  golang.MongoDBDatabase {
		_, err = file.WriteString(contentMongoDB)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeTokenModelFile(projectName, projectDatabase string){
	var contentMongoDB = `package models

// User model used to receive request and send JSON response, is used on authentication controller.
type Token struct {
	CreatedAt string ` + "`" + `json:"created_at,omitempty" bson:"created_at,omitempty"` + "`" + `
	Token     string ` + "`" + `json:"token,omitempty" bson:"token,omitempty"` + "`" + `
}
`
	var contentFirebase = `package models

// User model used to receive request and send JSON response, is used on authentication controller.
type Token struct {
	CreatedAt string ` + "`" + `json:"created_at,omitempty"` + "`" + `
	Token     string ` + "`" + `json:"token,omitempty"` + "`" + `
}
`
	file, err := os.Create(projectName + "/api/models/Token.go")
	if err != nil {
		log.Fatal(err)
	}

	if projectDatabase == golang.NoSelection || projectDatabase == golang.FirebaseDatabase {
		_, err = file.WriteString(contentFirebase)
		if err != nil {
			log.Fatal(err)
		}
	}

	if projectDatabase ==  golang.MongoDBDatabase {
		_, err = file.WriteString(contentMongoDB)
		if err != nil {
			log.Fatal(err)
		}
	}
}