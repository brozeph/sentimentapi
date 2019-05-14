package version

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

type Response struct {
	Build   string `json:"build"`
	Package string `json:"package"`
	Version string `json:"version"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getVersion)
	return router
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Build:   Build,
		Package: Package,
		Version: Version,
	}
	render.JSON(w, r, response)
}
