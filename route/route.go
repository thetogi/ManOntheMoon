package route

import (
	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {

	//TODO paginate data responses for requests for large data sources
	return InitRoutes()
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetApplicationRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
