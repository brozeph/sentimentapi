package main

import (
	"fmt"
	"log"
	"net/http"

	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/middleware"
	render "github.com/go-chi/render"

	version "github.com/brozeph/sentimentapi/resources"
	settings "github.com/brozeph/sentimentapi/settings"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-type headers as application/json
		middleware.Logger,          // Log API requests
		middleware.DefaultCompress, // gzip JSON responses
		middleware.RedirectSlashes, // redirect slaches to no slash URL version
		middleware.Recoverer,       // Recover from panic without server crash
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/version", version.Routes())
	})

	return router
}

func main() {
	settings, err := settings.New()
	if err != nil {
		log.Panicln("exception occurred loading settings", err)
	}

	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	// TODO: build in settings
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", settings.Server.Port), router))
}
