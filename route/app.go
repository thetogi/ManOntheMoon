package route

import (
	"ManOnTheMoonReviewService/controllers"
	"ManOnTheMoonReviewService/controllers/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func SetApplicationRoutes(router *mux.Router) *mux.Router {

	apiRouterV1 := router.PathPrefix("/api/v1").Subrouter()
	apiRouterV1Secure := router.PathPrefix("/api/v1/").Subrouter()
	apiRouterV2 := router.PathPrefix("/api/v2/").Subrouter()

	//Application routes
	ac := controllers.AppController{}
	router.NotFoundHandler = http.HandlerFunc(ac.NotFound)
	apiRouterV1.HandleFunc("/", ac.Home).Methods("GET")
	apiRouterV1.HandleFunc("/Health-Check", ac.HealthCheck).Methods("GET")

	//Player routes
	pc := controllers.PlayerController{}
	apiRouterV1.HandleFunc("/Player", pc.GetPlayer).Methods("GET")
	apiRouterV1.HandleFunc("/Player", pc.CreatePlayer).Methods("POST")
	apiRouterV1.HandleFunc("/Players/", pc.GetAllPlayers).Methods("GET")

	//Session routes
	sc := controllers.SessionController{}
	apiRouterV1.HandleFunc("/Session", sc.GetSession).Methods("GET")
	apiRouterV1.HandleFunc("/Session", sc.CreateSession).Methods("POST")
	apiRouterV1.HandleFunc("/Sessions/", sc.GetAllSessions).Methods("GET")

	//Rating routes
	rc := controllers.RatingController{}
	apiRouterV1.HandleFunc("/Random/Rating", rc.GetRandomRating).Methods("GET")
	apiRouterV1.HandleFunc("/Rating", rc.CreateRating).Methods("POST")

	apiRouterV2.HandleFunc("/Random/Rating", rc.GetRandomRating2).Methods("GET")

	//Secured routes
	auc := auth.AuthenticationController{}
	apiRouterV1Secure.Use(auc.JWTValidateMiddleware)
	apiRouterV1Secure.HandleFunc("/Rating", rc.GetRating).Methods("GET")
	apiRouterV1Secure.HandleFunc("/Ratings/", rc.GetRatings).Methods("GET")

	return router
}
