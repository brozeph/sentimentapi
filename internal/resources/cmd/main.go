package resources

import (
	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/middleware"
	render "github.com/go-chi/render"

	persons "github.com/brozeph/sentimentapi/internal/resources/persons"
	version "github.com/brozeph/sentimentapi/internal/resources/version"
	settings "github.com/brozeph/sentimentapi/internal/settings"
)

func Routes(settings *settings.Settings) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-type headers as application/json
		middleware.Logger,          // Log API requests
		middleware.DefaultCompress, // gzip JSON responses
		middleware.RedirectSlashes, // redirect slaches to no slash URL version
		middleware.Recoverer,       // Recover from panic without server crash
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/persons", persons.Routes())
		r.Mount("/version", version.Routes())
	})

	return router
}

func main() {
}
