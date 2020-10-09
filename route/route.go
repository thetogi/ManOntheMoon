package route

import (
	"ManOnTheMoonReviewService/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", controllers.Home).Methods("GET")
	routes.HandleFunc("/toggleShowEnv", controllers.ShowEnvToggle)

	routes.HandleFunc("/GameSession/Rating", controllers.ShowEnvToggle).Methods("POST")
	routes.HandleFunc("/GameSession/Rating", controllers.ShowEnvToggle).Methods("POST")
	routes.HandleFunc("/GameSession/Rating", controllers.GetRatingRandom).Methods("GET")
	routes.HandleFunc("/GameSession/{sessionId}", controllers.ShowEnvToggle)

	//Game routes
	routes.HandleFunc("/Game/{GameId}", controllers.GetGameByIdSql).Methods("GET")
	routes.HandleFunc("/Game/Create", controllers.PostGameCreateSql).Methods("POST")

	//Player routes
	routes.HandleFunc("/Game/{GameId}/Player/{PlayerId}", controllers.GetPlayerByIdSql).Methods("GET")
	routes.HandleFunc("/Game/{GameId}/Player/Create", controllers.PostPlayerCreateSql).Methods("POST")
	routes.HandleFunc("/Game/{GameId}/Players/", controllers.GetAllPlayersSql).Methods("GET")

	//Session routes
	routes.HandleFunc("/Session/{SessionId}", controllers.GetSessionByIdSql).Methods("GET")
	routes.HandleFunc("/Session/Create", controllers.PostSessionCreateSql).Methods("POST")
	routes.HandleFunc("/Sessions/", controllers.GetAllSessionsSql).Methods("GET")

	//Session Ratings routes
	routes.HandleFunc("/Session/{SessionId}/Rating", controllers.GetSessionRatingBySessionIdSql).Methods("GET")
	routes.HandleFunc("/Session/{SessionId}/CreateRating", controllers.PostSessionRatingCreateSql).Methods("POST")
	routes.HandleFunc("/Session/Ratings/", controllers.GetAllSessionRatingsSql).Methods("GET")
	return routes
}
