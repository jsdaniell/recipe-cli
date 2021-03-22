package go_content_files

import (
	"log"
	"os"
)

func CreateUtilsPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/api/utils", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(projectName + "/api/utils/json_utility", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(projectName + "/api/utils/security", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeJSONUtilityFile(projectName)
	writeTokenFile(username, projectName)
}

func writeJSONUtilityFile(projectName string){
	var content = `package json_utility

import (
	"encoding/json"
)

// Utility function to lower case JSON pure structs, generally used on response of a request inside controllers.
func StructToLowerCaseJson(data interface{}) (interface{}, error){

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var r interface{}

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
`

	file, err := os.Create(projectName + "/api/utils/json_utility/json_utility.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeTokenFile(username, projectName string){
	var content = `package security

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/`+ username +`/`+ projectName +`/api/models"
	"time"
)

const layout = "2006-01-02T15:04:05.000Z"

// CreateToken Creates the token stored on each user, receives a generic string to generate token based on string.
func CreateToken(s string) models.Token {
	h := sha1.New()
	h.Write([]byte(s))

	hashes := hex.EncodeToString(h.Sum(nil))
	actual := time.Now()
	date := actual.Format(time.RFC3339)

	var t = models.Token{
		CreatedAt: date,
		Token:     hashes,
	}

	return t
}

func ValidateToken(tk models.Token) error {

	t, err := time.Parse(time.RFC3339, tk.CreatedAt)
	if err != nil {
		return err
	}

	dayLater := t.Add(24 * time.Hour)

	if !t.Before(dayLater) {
		return errors.New("token expired")
	}

	return nil
}
`

	file, err := os.Create(projectName + "/api/utils/security/token.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}