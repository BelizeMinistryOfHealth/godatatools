package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// FindUserByID returns the user that has the corresponding ID
func (s *Store) FindUserByID(ctx context.Context, ID string) (*models.User, error) {
	col := s.Client.Database(s.Database).Collection(userCollection)
	filter := bson.D{{"_id", ID}} //nolint:govet
	var user models.User
	err := col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, MongoQueryErr{
			Reason: "FindUserByID() error",
			Inner:  err,
		}
	}

	return &user, err
}
