package go_content_files

import (
	"log"
	"os"
)

func CreateResponsesPackage(projectName string) {


	err := os.Mkdir(projectName + "/api/responses", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeResponsesFile(projectName)
}

func writeResponsesFile(projectName string) {
	var content = `// Response utility to returns JSON error or JSON response to the request.
package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string` + "`" + `json:"error"` + "`" + `
}{
Error: err.Error(),
})
return
}

JSON(w, http.StatusBadRequest, nil)
}
`

	file, err := os.Create(projectName + "/api/responses/responses.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}