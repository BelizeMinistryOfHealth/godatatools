package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	mn "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Store struct {
	Database   string
	URI        string
	Client     *mn.Client
	Connect    func(context.Context) error
	Disconnect func(ctx context.Context) error
}

const personCollection = "person"
const outbreakCollection = "outbreak"
const labCollection = "labResult"

func New(uri, database string) (Store, error) {
	client, err := mn.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return Store{}, MongoConnectionErr{
			Reason: "failed to create mongo client",
			Inner:  err,
		}
	}
	return Store{
		Database:   database,
		URI:        uri,
		Client:     client,
		Connect:    client.Connect,
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

// FindCasesByReportingDate returns the cases reported in the specified date for the designated outbreak.
func (s *Store) FindCasesByReportingDate(ctx context.Context, outbreakID string, reportingDate time.Time) ([]models.Case, error) {
	collection := s.Client.Database(s.Database).Collection(personCollection)
	filter := bson.M{"outbreakId": outbreakID, "dateOfReporting": reportingDate}
	var cases []models.Case
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return cases, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error querying for cases on reporting date: %v for outbreak: %s", reportingDate, outbreakID),
			Inner:  err,
		}
	}

	if err := cursor.All(ctx, &cases); err != nil {
		return cases, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error decoding cases for reporting date: %v and outbreakId: %s", reportingDate, outbreakID),
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
			Inner:  err,
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

// OutbreakById returns an outbreak that has the specified id
func (s *Store) OutbreakById(ctx context.Context, ID string) (models.Outbreak, error) {
	collection := s.Client.Database(s.Database).Collection(outbreakCollection)
	var outbreak models.Outbreak
	filter := bson.M{"_id": ID}
	result := collection.FindOne(ctx, filter)
	if result == nil {
		return outbreak, MongoNoResultErr{
			Reason: fmt.Sprintf("no outbreak found with id: %s", ID),
			Inner:  nil,
		}
	}
	if err := result.Decode(&outbreak); err != nil {
		return outbreak, MongoQueryErr{
			Reason: fmt.Sprintf("mongo: could not decode outbreak with id: %s", ID),
			Inner:  err,
		}
	}

	return outbreak, nil
}

// LabTestsForCases returns a list of RawLabTest for cases provided in the caseIds
func (s *Store) LabTestsForCases(ctx context.Context, caseIds []string) ([]models.RawLabTest, error) {
	collection := s.Client.Database(s.Database).Collection(labCollection)
	filter := bson.M{"personId": bson.M{"$in": caseIds}, "deleted": false}
	var labTests []models.RawLabTest
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error querying lab tests for caseIds %v", caseIds),
			Inner:  err,
		}
	}

	if err := cursor.All(ctx, &labTests); err != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error decoding lab tests for caseIds: %v", caseIds),
			Inner:  err,
		}
	}
	return labTests, nil
}
