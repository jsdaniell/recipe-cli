package golang_socket_content_files

import (
	"log"
	"os"
)

func CreateMainFile(username, projectName string) {
	var content = `package main

import (
	"github.com/ambelovsky/gosf"
	"github.com/`+ username +`/`+ projectName +`/router"
)


func init() {
	router.RegisterRoutes() // Configure endpoint request handlers

	// Load Config File Based on Environmental Configuration
	if value, exist := gosf.App.Env["GOSF_ENV"]; exist && value != "dev" {
		// Prod/Stage Config
		gosf.LoadConfig("server", "server-secure.json")
	} else {
		// Default and "dev" config
		gosf.LoadConfig("server", "server.json")
	}
}

func main() {
	// Start the server
	serverConfig := gosf.App.Config["server"].(map[string]interface{})
	gosf.Startup(serverConfig)
}`

	file, err := os.Create(projectName + "/main.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateHTMLFile(projectName string) {
	var content = `<!DOCTYPE html>
<html>
<head>
    <title>Socket</title>
</head>
<body>
<script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.2.0/socket.io.slim.js"></script>
<script>
    var socket = io.connect('ws://localhost:9999', { transports: ['websocket'] });

    socket.emit('echo', { text: 'Hello world.' }, function(response) {
        console.log(response);
    });
</script>
</body>
</html>`

	file, err := os.Create(projectName + "/index.html")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}