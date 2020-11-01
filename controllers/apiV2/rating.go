package apiV2

import (
	"ManOnTheMoonReviewService/controllers"
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db/seed/seeder"
	"net/http"
)

type RatingController struct {
	controllers.Controller
}

//GetRandomRating Simulates returning a rating. Nothing is retrieved from or committed to database.
func (r *RatingController) GetRandomRating(w http.ResponseWriter, req *http.Request) {
	response.Write(w, response.Response{
		Code: http.StatusOK,
		Data: seed.MockRatingData(),
	})
}
