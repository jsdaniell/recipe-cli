package routes

import (
	"github.com/jsdaniell/fire/api/controllers"
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
