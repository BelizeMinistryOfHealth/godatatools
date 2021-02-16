package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	mn "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Database string
	URI string
	Client *mn.Client
	Connect func(context.Context) error
	Disconnect func(ctx context.Context) error
}

const personCollection = "person"
const outbreakCollection = "outbreak"

func New(uri, database string) (Store, error) {
	client, err := mn.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return Store{}, MongoConnectionErr{
			Reason: "failed to create mongo client",
			Inner:  err,
		}
	}
	return Store{
		Database: database,
		URI: uri,
		Client: client,
		Connect: client.Connect,
		Disconnect: client.Disconnect,
	}, nil
}

func (s *Store) FindCasesByOutbreak(ctx context.Context, ID string) ([]models.Case, error) {
	collection := s.Client.Database(s.Database).Collection(personCollection)
	filter := bson.M{"outbreakId": ID}
	var cases []models.Case
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return cases, MongoQueryErr{
			Reason: "error searching for cases by outbreakId",
			Inner:  err,
		}
	}

	if err := cursor.All(ctx, &cases); err != nil {
		return cases, MongoQueryErr{
			Reason: fmt.Sprintf("failed to decode the result of cases for the outbreak: %s", ID),
			Inner:  err,
		}
	}
	return cases, nil
}


func (s *Store) ListOutbreaks(ctx context.Context) ([]models.Outbreak, error) {
	collection := s.Client.Database(s.Database).Collection(outbreakCollection)
	var outbreaks []models.Outbreak
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return outbreaks, MongoQueryErr{
			Reason: "error listing outbreaks",
			Inner: err,
		}
	}
	if err := cursor.All(ctx, &outbreaks); err != nil {
		return outbreaks, MongoQueryErr{
			Reason: "failed to decode the list of outbreaks",
			Inner:  err,
		}
	}
	return outbreaks, nil
}
