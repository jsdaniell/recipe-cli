package go_content_files

import (
	"log"
	"os"
)

func CreateDbPackage(projectName, projectDatabase string) {
	err := os.Mkdir(projectName + "/api/db", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeDatabaseFile(projectName, projectDatabase)
}

func writeDatabaseFile(projectName, projectDatabase string) {
	var contentFirebase = `// Package of client configuration of firestore and activation
package db

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

// Configuration file setting firebase credentials and returning the client to use on database operations
// localized on repository package.
func FirestoreClient() *firestore.Client{
	sa := option.WithCredentialsFile("entity-firebase-adminsdk.json")

	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}



	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
`
	var contentMongoDB = `package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

var EntitiesCollection = "entities"

//I have used below constants just to hold required database config.
const (
	CONNECTIONSTRING = "mongodb+srv://${username}:connection_string*"
	DB               = "database_name"
)

// Configuration file setting firebase credentials and returning the client to use on database operations
// localized on repository package.
func ConnectDB() (*mongo.Database, error) {

	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		fmt.Println("Connected to MongoDB")

		clientInstance = client
	})
	return clientInstance.Database(DB), clientInstanceError

}
`

	var contentNoSelection = `package db

// Put your database connection here.
`

	file, err := os.Create(projectName + "/api/db/database.go")
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