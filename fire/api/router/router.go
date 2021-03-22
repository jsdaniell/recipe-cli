// Router setup.
package router

import (
	"github.com/gorilla/mux"
	"github.com/jsdaniell/fire/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRoutesWithMiddlewares(r)
}
