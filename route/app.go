package route

import (
	"ManOnTheMoonReviewService/controllers"
	"ManOnTheMoonReviewService/controllers/apiV1"
	"ManOnTheMoonReviewService/controllers/apiV2"
	"ManOnTheMoonReviewService/controllers/auth"
	"github.com/gorilla/mux"
)

func SetApplicationRoutes(router *mux.Router) *mux.Router {

	ac := controllers.AppController{}
	pc := apiV1.PlayerController{}
	sc := apiV1.SessionController{}
	auc := auth.AuthenticationController{}

	apiRouterV1 := router.PathPrefix("/api/v1").Subrouter()
	apiRouterV2 := router.PathPrefix("/api/v2").Subrouter()

	rc := apiV1.RatingController{}
	rc2 := apiV2.RatingController{}

	//Application routes
	apiRouterV1.HandleFunc("/", ac.Home).Methods("GET")
	apiRouterV1.HandleFunc("/Health-Check", ac.HealthCheck).Methods("GET")
	apiRouterV1.HandleFunc("/Random/Rating", rc.GetRandomRating).Methods("GET")
	apiRouterV2.HandleFunc("/Random/Rating", rc2.GetRandomRating).Methods("GET")
	//Player routes
	apiRouterV1.HandleFunc("/Player", pc.GetPlayer).Methods("GET")
	apiRouterV1.HandleFunc("/Player", pc.CreatePlayer).Methods("POST")
	apiRouterV1.HandleFunc("/Players/", pc.GetAllPlayers).Methods("GET")

	//Session routes
	apiRouterV1.HandleFunc("/Session", sc.GetSession).Methods("GET")
	apiRouterV1.HandleFunc("/Session", sc.CreateSession).Methods("POST")
	apiRouterV1.HandleFunc("/Sessions/", sc.GetAllSessions).Methods("GET")

	//Rating routes
	apiRouterV1.Use(auc.JWTValidateMiddleware)
	apiRouterV1.HandleFunc("/Rating", rc.GetRating).Methods("GET")
	apiRouterV1.HandleFunc("/Rating", rc.CreateRating).Methods("POST")
	apiRouterV1.HandleFunc("/Ratings/", rc.GetRatings).Methods("GET")

	return router
}
