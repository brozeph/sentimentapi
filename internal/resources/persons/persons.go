package persons

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	Build   string
	Package string
	Version string
)

type PersonResponse struct {
	PersonId string `json:"personId"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{personId}", getPerson)
	return router
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	personId := chi.URLParam(r, "personId")

	response := PersonResponse{
		PersonId: personId,
		Email:    "test@test.com",
		Name:     "test",
	}
	render.JSON(w, r, response)
}
