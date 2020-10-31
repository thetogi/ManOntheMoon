package route

import (
	"ManOnTheMoonReviewService/controllers"
	"github.com/gorilla/mux"
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

	//Application routes
	routes.HandleFunc("/", ac.Home).Methods("GET")
	routes.HandleFunc("/api/v1/Health-Check", ac.HealthCheck).Methods("GET")
	routes.HandleFunc("/api/v1/Random/Rating", rc.GetRandomRating).Methods("GET")

	//Player routes
	routes.HandleFunc("/api/v1/Player/{PlayerId}", pc.GetPlayer).Methods("GET")
	routes.HandleFunc("/api/v1/Player/Create", pc.CreatePlayer).Methods("POST")
	routes.HandleFunc("/api/v1/Players/", pc.GetAllPlayers).Methods("GET")

	//Session routes
	routes.HandleFunc("/api/v1/Session", sc.GetSession).Methods("GET")
	routes.HandleFunc("/api/v1/Session", sc.CreateSession).Methods("POST")
	routes.HandleFunc("/api/v1/Sessions/", sc.GetAllSessions).Methods("GET")

	//Rating routes
	routes.HandleFunc("/api/v1/Rating", rc.GetRating).Methods("GET")
	routes.HandleFunc("/api/v1/Rating", rc.CreateRating).Methods("POST")
	routes.HandleFunc("/api/v1/Ratings/", rc.GetRatings).Methods("GET")
	//TODO paginate data responses for requests for large data sources
	return routes
}
