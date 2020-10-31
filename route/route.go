package route

import (
	"ManOnTheMoonReviewService/controllers"
	"github.com/gorilla/mux"
	"os"
)

func GetRoutes() *mux.Router {
	//Define RESTful API Endpoints for service

	//Define router to associate routes/URIs
	routes := mux.NewRouter()
	ac := controllers.AppController{}
	pc := controllers.PlayerController{}
	sc := controllers.SessionController{}
	rc := controllers.RatingController{}

	//TODO Structure routes in loop and incorporate middleware handling

	apiVer := os.Getenv("API_VER")

	if apiVer == "" {
		panic("API version not configured")
	}

	apiBasePath := "/api/" + apiVer

	//Application routes
	routes.HandleFunc("/", ac.Home).Methods("GET")
	routes.HandleFunc(apiBasePath+"/Health-Check", ac.HealthCheck).Methods("GET")
	routes.HandleFunc(apiBasePath+"/Random/Rating", rc.GetRandomRating).Methods("GET")

	//Player routes
	routes.HandleFunc(apiBasePath+"/Player", pc.GetPlayer).Methods("GET")
	routes.HandleFunc(apiBasePath+"/Player", pc.CreatePlayer).Methods("POST")
	routes.HandleFunc(apiBasePath+"/Players/", pc.GetAllPlayers).Methods("GET")

	//Session routes
	routes.HandleFunc(apiBasePath+"/Session", sc.GetSession).Methods("GET")
	routes.HandleFunc(apiBasePath+"/Session", sc.CreateSession).Methods("POST")
	routes.HandleFunc(apiBasePath+"/Sessions/", sc.GetAllSessions).Methods("GET")

	//Rating routes
	routes.HandleFunc(apiBasePath+"/Rating", rc.GetRating).Methods("GET")
	routes.HandleFunc(apiBasePath+"/Rating", rc.CreateRating).Methods("POST")
	routes.HandleFunc(apiBasePath+"/Ratings/", rc.GetRatings).Methods("GET")
	//TODO paginate data responses for requests for large data sources
	return routes
}
