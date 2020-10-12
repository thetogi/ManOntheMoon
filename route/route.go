package route

import (
	"ManOnTheMoonReviewService/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	//Define RESTful API Endpoints for service

	//Define router to associate routes/URIs
	routes := mux.NewRouter()

	//TODO Structure routes in loop and incorporate middleware handling

	//Internal Tools Routes
	routes.HandleFunc("/", controllers.Home).Methods("GET")
	routes.HandleFunc("/Health-Check", controllers.HealthCheck).Methods("GET")
	routes.HandleFunc("/GameSession/Rating", controllers.GetRatingRandomForTesting).Methods("GET")

	//Player routes
	routes.HandleFunc("/Player/{PlayerId}", controllers.GetPlayerByIdSql).Methods("GET")
	routes.HandleFunc("/Player/Create", controllers.PostPlayerCreateSql).Methods("POST")
	routes.HandleFunc("/Players/", controllers.GetAllPlayersSql).Methods("GET")

	//Session routes
	routes.HandleFunc("/Session/{SessionId}", controllers.GetSessionByIdSql).Methods("GET")
	routes.HandleFunc("/Session/Create", controllers.PostSessionCreateSql).Methods("POST")
	routes.HandleFunc("/Sessions/", controllers.GetAllSessionsSql).Methods("GET")

	//Session API En
	routes.HandleFunc("/Session/{SessionId}/Rating", controllers.GetSessionRatingBySessionIdSql).Methods("GET")
	routes.HandleFunc("/Session/{SessionId}/CreateRating", controllers.PostSessionRatingCreateSql).Methods("POST")
	routes.HandleFunc("/Session/Ratings/", controllers.GetAllSessionRatingsSql).Methods("GET")
	//TODO paginate data responses for requests for large data sources
	return routes
}
