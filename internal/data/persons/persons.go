package persons

import (
	"context"
	"log"

	"github.com/brozeph/sentimentapi/internal/settings"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
type Person struct {
	PersonID string
	Email    string
	Name     string
}
*/

type PersonMapper struct {
	client     *mongo.Client
	collection *mongo.Collection
	context    context.Context
}

func NewPersonMapper(settings *settings.Settings, client *mongo.Client, ctx context.Context) *PersonMapper {
	m := new(PersonMapper)

	m.client = client
	m.collection = client.Database(settings.Data.Mongo.Database).Collection("persons")
	m.context = ctx

	return m
}

func (m PersonMapper) Insert(person *interface{}) error {
	result, err := m.collection.InsertOne(m.context, person)
	if err != nil {
		log.Printf("error encountered inserting person: %s", err)
	} else {
		log.Printf("successfully inserted person with insert ID %s", result.InsertedID)
	}

	return err
}
