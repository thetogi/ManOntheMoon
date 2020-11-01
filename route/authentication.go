package route

import (
	"ManOnTheMoonReviewService/controllers/auth"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	ac := auth.AuthenticationController{}
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/authenticate", ac.Login).Methods("POST")
	authRouter.HandleFunc("/logout", ac.Logout).Methods("POST")
	return router
}
