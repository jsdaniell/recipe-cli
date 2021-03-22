package routes

import (
	"github.com/jsdaniell/fire/api/controllers"
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
