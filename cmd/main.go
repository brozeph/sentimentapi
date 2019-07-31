package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	chi "github.com/go-chi/chi"

	data "github.com/brozeph/sentimentapi/internal/data/cmd"
	resources "github.com/brozeph/sentimentapi/internal/resources/cmd"
	settings "github.com/brozeph/sentimentapi/internal/settings"
)

func main() {
	settings, err := settings.New()
	if err != nil {
		log.Panicln("exception occurred loading settings", err)
	}

	timeout, err := time.ParseDuration(settings.Data.Mongo.Timeout)
	if err != nil {

	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	data, err := data.NewDataClient(settings, ctx)
	if err != nil {
		log.Panicln("exception occurred establishing connection to database", err)
	}

	// check connection to Mongodb
	if err := data.Client.Ping(ctx, nil); err != nil {
		log.Panicln("unable to ping MongoDB")
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
