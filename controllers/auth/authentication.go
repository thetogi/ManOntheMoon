package auth

import (
	"ManOnTheMoonReviewService/controllers"
	"ManOnTheMoonReviewService/controllers/response"
	auth "ManOnTheMoonReviewService/models/auth"
	"ManOnTheMoonReviewService/util"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	gorillaContext "github.com/gorilla/context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthenticationController struct {
	controllers.Controller
	User auth.User
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func (a *AuthenticationController) Login(w http.ResponseWriter, req *http.Request) {
	var credentials auth.User
	err := util.ParseRequestBody(w, req, &credentials)
	if err != nil {
		return
	}

	// Check in your db if the user exists or not
	if credentials.Username == "jon" && credentials.Password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		secret := os.Getenv("JWT_SECRET")

		// Set claims. This information could be used for additional validation of user roles, etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID works too)
		t, err := token.SignedString([]byte(secret))
		if err != nil {
			return
		}
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: map[string]string{"token": t},
		})
		return
	}
	response.Write(w, response.Response{
		Code: http.StatusUnauthorized,
	},
	)
}

func (a *AuthenticationController) Logout(w http.ResponseWriter, req *http.Request) {
	authorizationHeader := req.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				secret := os.Getenv("JWT_SECRET")
				return []byte(secret), nil
			})
			if error != nil {
				response.Write(w, response.Response{
					Code:    http.StatusBadRequest,
					Action:  "Authentication",
					Message: "Error parsing token",
					Errors:  map[string]string{"Error": error.Error()},
				})
				return
			}
			if token.Valid {
				err := a.User.Logout(req.Context(), token)
				log.Println(token.Raw)
				if err != nil {
					response.Write(w, response.Response{
						Code:    http.StatusBadRequest,
						Action:  "LogOut",
						Message: "An error ocurred while logging out",
						Errors:  map[string]string{"Error": err.Error()},
					})
				}
				response.Write(w, response.Response{
					Code: http.StatusOK,
					Data: "Log out successful",
				})
			} else {
				response.Write(w, response.Response{
					Code:    http.StatusBadRequest,
					Action:  "Authentication",
					Message: "Invalid authorization token",
				})
				return
			}
		}
	} else {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "Authentication",
			Message: "An authorization header is required",
		})
	}
}

func (a *AuthenticationController) JWTValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("parsing header authorization")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				log.Println("splitting authorization")
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					secret := os.Getenv("JWT_SECRET")
					return []byte(secret), nil
				})
				if error != nil {
					v, _ := error.(*jwt.ValidationError)
					if v.Errors == jwt.ValidationErrorExpired {
						//Todo: Implement refresh logic if desired.
					}

					response.Write(w, response.Response{
						Code:    http.StatusBadRequest,
						Action:  "Authentication",
						Message: "Error parsing token",
						Errors:  map[string]string{"Error": error.Error()},
					})
					return
				}
				log.Println("checking if valid")
				if a.isTokenValid(r.Context(), token) {
					gorillaContext.Set(r, "decoded", token.Claims)
					next.ServeHTTP(w, r)
				} else {
					response.Write(w, response.Response{
						Code:    http.StatusBadRequest,
						Action:  "Authentication",
						Message: "Invalid authorization token",
					})
					return
				}
			}
			log.Println("wierd spot")
		} else {

			response.Write(w, response.Response{
				Code:    http.StatusBadRequest,
				Action:  "Authentication",
				Message: "An authorization header is required",
			})
		}
	})
}

func (a *AuthenticationController) isTokenValid(c context.Context, token *jwt.Token) bool {
	if !token.Valid || a.User.IsTokenBlacklisted(c, token.Raw) {
		return false
	}
	return true
}
