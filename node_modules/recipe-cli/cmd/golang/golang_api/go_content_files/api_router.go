package go_content_files

import (
	"log"
	"os"
)

func CreateRouterPackage(username, projectName string) {
	err := os.Mkdir(projectName + "/api/router", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName + "/api/router/routes", os.FileMode(0777))
	if err != nil {
		log.Fatal(err)
	}

	writeRoutesFile(username, projectName)
	writeServerRoutesFile(username, projectName)
	writeEntityRoutesFile(username, projectName)
	writeRouterFile(username, projectName)
}

func writeRoutesFile(username, projectName string){
	var content = `package routes

import (
	"github.com/gorilla/mux"
	"github.com/`+ username +`/`+ projectName +`/api/middlewares"
	"net/http"
)

// Route struct provides the structure of route
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	Open    bool
}

// Load all routes and setting with middlewares to receive requests.
func Load() []Route {

	routes := [][]Route{
		serverRoutes,
		entityRoutes,
	}

	var joinedRoutes []Route

	for _, r := range routes {
		joinedRoutes = append(joinedRoutes, r...)
	}

	return joinedRoutes
}

// SetupRoutesWithMiddlewares Using middlewares on handle of route.
func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.URI, middlewares.SetMiddlewareLogger(
			middlewares.SetMiddlewareJSON(route.Handler, route.Open))).Methods(route.Method, "OPTIONS")
	}

	return r
}
`

	file, err := os.Create(projectName + "/api/router/routes/routes.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeServerRoutesFile(username, projectName string){
	var content = `package routes

import (
	"github.com/`+ username +`/`+ projectName +`/api/controllers"
	"net/http"
)

// Generic server routes with their controller.
var serverRoutes = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: controllers.ServerRunning,
		Open: true,
	},
}
`

	file, err := os.Create(projectName + "/api/router/routes/server_routes.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeEntityRoutesFile(username, projectName string){
	var content = `package routes

import (
	"github.com/`+ username +`/`+ projectName +`/api/controllers"
	"net/http"
)

var entityRoutes = []Route{
	{
		URI:     "/addEntity",
		Method:  http.MethodPost,
		Handler: controllers.AddEntityController,
		Open:    true,
	},
	// Route{
	// 	Uri:     "/deleteEntity/{code}",
	// 	Method:  http.MethodDelete,
	// 	Handler: controllers.DeleteProductController,
	// 	Open:    true,
	// },
}
`

	file, err := os.Create(projectName + "/api/router/routes/entity_routes.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func writeRouterFile(username, projectName string){
	var content = `// Router setup.
package router

import (
	"github.com/gorilla/mux"
	"github.com/`+ username +`/`+ projectName +`/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRoutesWithMiddlewares(r)
}
`

	file, err := os.Create(projectName + "/api/router/router.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}