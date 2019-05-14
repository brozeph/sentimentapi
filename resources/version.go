package version

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getVersion)
	return router
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	version := Version{
		Name:    "sentimentapi",
		Version: "v1.0.0",
	}
	render.JSON(w, r, version)
}
