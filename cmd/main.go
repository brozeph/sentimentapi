package main

import (
	"fmt"
	"log"
	"net/http"

	chi "github.com/go-chi/chi"

	resources "github.com/brozeph/sentimentapi/internal/resources/cmd"
	settings "github.com/brozeph/sentimentapi/internal/settings"
)

func main() {
	settings, err := settings.New()
	if err != nil {
		log.Panicln("exception occurred loading settings", err)
	}

	router := resources.Routes(settings)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", settings.Server.Port), router))
}
