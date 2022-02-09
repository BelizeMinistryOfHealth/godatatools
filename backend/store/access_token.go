package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// FindUserIDForAccessToken returns the userId that belongs to the access token
func (s *Store) FindUserIDForAccessToken(ctx context.Context, accessToken string) (string, error) {
	col := s.Client.Database(s.Database).Collection(accessTokenCollection)
	filter := bson.M{
		"$and": bson.A{
			bson.M{"_id": accessToken},
			bson.M{"deleted": false},
		},
	}
	var token models.AccessToken
	err := col.FindOne(ctx, filter).Decode(&token)
	if err != nil {
		return "", MongoQueryErr{
			Reason: "FindUserIdForAccessToken() error",
			Inner:  err,
		}
	}
	return token.UserID, nil
}
