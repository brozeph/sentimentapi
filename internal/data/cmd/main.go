package data

import (
	"context"
	"log"

	"github.com/brozeph/sentimentapi/internal/data/persons"
	"github.com/brozeph/sentimentapi/internal/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client struct is composed of the mongo client, context, settings and a mapper for
// translating person objects to the underlying data store.
type Client struct {
	Client       *mongo.Client
	Context      context.Context
	PersonMapper *persons.PersonMapper
	Settings     *settings.Settings
}

func NewClient(settings *settings.Settings, ctx context.Context) (*Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(settings.Data.Mongo.URL))
	if err != nil {
		log.Printf("unable to create MongoDB client with URL: %s", settings.Data.Mongo.URL)
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		log.Printf(
			"unable to connect to MongoDB at %s in %s",
			settings.Data.Mongo.URL,
			settings.Data.Mongo.Timeout)
		return nil, err
	}

	// construct new data client
	dataClient := new(Client)
	dataClient.Client = client
	dataClient.Context = ctx
	dataClient.PersonMapper = persons.NewPersonMapper(settings, client, ctx)
	dataClient.Settings = settings

	return dataClient, nil
}

func main() {}
