package golang_socket_content_files

import (
	"log"
	"os"
)

func CreateConfigPackage(projectName string) {
	err := os.Mkdir(projectName + "/config", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeServerJSON(projectName)
	writeServerSecurityJSON(projectName)
}

func writeServerJSON(projectName string){
	var content = `{
  "port": 9990,
  "path": "/"
}
`

	file, err := os.Create(projectName + "/config/server.json")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeServerSecurityJSON(projectName string){
	var content = `{
  "port": 9999,
  "path": "/",
  "secure": true,
  "ssl-cert": "ssl.crt",
  "ssl-key": "ssl.key"
}
`

	file, err := os.Create(projectName + "/config/server-secure.json")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

