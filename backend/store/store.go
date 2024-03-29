package store

import (
	"bz.moh.epi/godatatools/age"
	"bz.moh.epi/godatatools/models"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
const accessTokenCollection = "accessToken"
const userCollection = "user"

//go:embed locs2.json
var rawLocations []byte

type CasesById struct {
	Id    string
	Cases []models.Case
}

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

func (s *Store) FindCasesByOutbreak(ctx context.Context, ID string, startDate, endDate *time.Time) ([]models.Case, error) {
	var cases []models.Case

	if startDate.After(*endDate) {
		return cases, fmt.Errorf("startDate must be before endDate")
	}
	collection := s.Client.Database(s.Database).Collection(personCollection)
	filter := bson.M{
		"outbreakId": ID,
		"$and": bson.A{
			bson.M{"createdAt": bson.M{
				"$gte": startDate}},
			bson.M{"createdAt": bson.M{
				"$lt": endDate.Add(time.Hour * 24)}},
		}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return cases, MongoQueryErr{
			Reason: "error searching for cases by outbreakId",
			Inner:  err,
		}
	}

	if err := cursor.All(ctx, &cases); err != nil {
		curr := cursor.ID()
		return cases, MongoQueryErr{
			Reason: fmt.Sprintf("failed to decode the result of cases for the outbreak: %s curr: %v", ID, curr),
			Inner:  err,
		}
	}

	var result []models.Case
	for _, c := range cases {
		if c.Questionnaire.BhisNumber != nil {
			c.Bhis = c.Questionnaire.BhisNumber[0].Value
		}
		result = append(result, c)
	}
	return result, nil
}

// FindCasesByPersonIds will find all the cases for the corresponding person ids.
func (s *Store) FindCasesByPersonIds(ctx context.Context, IDs []string) (map[string]*models.LabTestCase, error) {
	var indexDB = make(map[string]*models.LabTestCase)
	collection := s.Client.Database(s.Database).Collection(personCollection)

	locations := make(map[string]models.AddressLocation)
	err := json.Unmarshal(rawLocations, &locations)
	if err != nil {
		return indexDB, MongoQueryErr{
			Reason: "failed to unmarshal locations",
			Inner:  err,
		}
	}

	filter := bson.M{
		"_id":     bson.M{"$in": IDs},
		"deleted": false,
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return indexDB, MongoQueryErr{
			Reason: "error searching for persons by ids",
			Inner:  err,
		}
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person models.LabTestCase
		if err = cursor.Decode(&person); err != nil {
			return indexDB, MongoQueryErr{
				Reason: "error decoding person from mongo",
				Inner:  err,
			}
		}
		// Find location
		addresses := person.Addresses
		if len(addresses) > 0 {
			//location, err := s.FindLocation(ctx, addresses[0].LocationId)
			location := locations[addresses[0].LocationId]
			if err == nil {
				person.Location = &location
			}
		}

		indexDB[person.ID] = &person
	}

	return indexDB, nil
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

// ListOutbreaks lists all the current outbreaks
func (s *Store) ListOutbreaks(ctx context.Context, outbreakIDS []string) ([]models.Outbreak, error) {
	collection := s.Client.Database(s.Database).Collection(outbreakCollection)
	var outbreaks []models.Outbreak
	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": outbreakIDS}})
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

// LabTestsByCaseName retrieves the lab tests for a case that matches the first & last names.
func (s *Store) LabTestsByCaseName(ctx context.Context, firstName, lastName string, outbreakIDs []string) ([]models.LabTest, error) {
	var rawLabTests []models.RawLabTest
	var labTests []models.LabTest
	/// find cases by first & last names.
	personCol := s.Client.Database(s.Database).Collection(personCollection)
	firstNameRegex := primitive.Regex{Pattern: firstName, Options: "i"}
	lastNameRegex := primitive.Regex{Pattern: lastName, Options: "i"}
	filter := bson.M{
		"firstName":  bson.M{"$regex": firstNameRegex},
		"lastName":   bson.M{"$regex": lastNameRegex},
		"outbreakId": bson.M{"$in": outbreakIDs},
	}
	var cases []models.Case
	cursor, err := personCol.Find(ctx, filter)
	if err != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error querying for cases records for %s %s", firstName, lastName),
			Inner:  err,
		}
	}
	if err := cursor.All(ctx, &cases); err != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("mongo error querying case records for %s %s", firstName, lastName),
			Inner:  err,
		}
	}
	// Retrieve the lab tests for every case
	var caseIds []string
	for _, c := range cases {
		caseIds = append(caseIds, c.ID)
	}

	if len(caseIds) == 0 {
		return labTests, nil
	}
	rawLabTests, testsErr := s.LabTestsForCases(ctx, caseIds)
	if testsErr != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("failed to fetch lab tests for %s %s", firstName, lastName),
			Inner:  testsErr,
		}
	}

	for i, _ := range rawLabTests { //nolint:gofmt,gosimple
		labTests = append(labTests, RawLabTestToLabTest(rawLabTests[i], findCaseInCases(rawLabTests[i].PersonId, cases)))
	}

	return labTests, nil
}

// RawLabTestToLabTest converts a RawLabTest to a LabTest
func RawLabTestToLabTest(test models.RawLabTest, person models.Case) models.LabTest {
	labFacility := models.LabFacilities[test.OutbreakID]

	status, err := models.ParseLabTestStatus(test.Status)
	if err != nil {
		status = "N/A"
	}

	testType, testTypeErr := models.ParseLabTestType(test.TestType)
	if testTypeErr != nil {
		testType = "N/A"
	}

	testResult, testResultErr := models.ParseLabTestResult(test.Result)
	if testResultErr != nil {
		testResult = "N/A"
	}
	personAge := age.AgeAt(person.Dob, time.Now())
	if test.DateTesting != nil {
		personAge = age.AgeAt(person.Dob, *test.DateTesting)
	}
	var gender = person.Gender
	if gender == "LNG_REFERENCE_DATA_CATEGORY_GENDER_MALE" {
		gender = "Male"
	}

	if gender == "LNG_REFERENCE_DATA_CATEGORY_GENDER_FEMALE" {
		gender = "Female"
	}

	labTest := models.LabTest{
		ID:                  test.ID,
		LabName:             labFacility.Name,
		PersonType:          test.PersonType,
		DateSampleTaken:     test.DateSampleTaken,
		DateSampleDelivered: test.DateSampleDelivered,
		DateTesting:         test.DateTesting,
		DateOfResult:        test.DateOfResult,
		SampleIdentifier:    test.SampleIdentifier,
		SampleType:          test.SampleType,
		TestType:            testType,
		Result:              testResult,
		Status:              status,
		OutbreakID:          test.OutbreakID,
		TestedFor:           test.TestedFor,
		CreatedAt:           test.CreatedAt,
		CreatedBy:           test.CreatedBy,
		UpdatedAt:           test.UpdatedAt,
		Person: models.Person{
			ID:        person.ID,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Gender:    gender,
			Dob:       person.Dob,
			Age:       personAge,
			Documents: person.Documents,
		},
		LabFacility: labFacility,
	}
	return labTest
}

// RawLabTestToLabReport converts a RawLabTest to a LabReport
func RawLabTestToLabReport(test models.RawLabTest, person models.LabTestCase) models.LabTestReport {
	labFacility := models.LabFacilities[test.OutbreakID]

	status, err := models.ParseLabTestStatus(test.Status)
	if err != nil {
		status = "N/A"
	}

	testType, testTypeErr := models.ParseLabTestType(test.TestType)
	if testTypeErr != nil {
		testType = "N/A"
	}

	testResult, testResultErr := models.ParseLabTestResult(test.Result)
	if testResultErr != nil {
		testResult = "N/A"
	}

	labTest := models.LabTestReport{
		ID:                  test.ID,
		LabName:             labFacility.Name,
		PersonType:          test.PersonType,
		DateSampleTaken:     test.DateSampleTaken,
		DateSampleDelivered: test.DateSampleDelivered,
		DateTesting:         test.DateTesting,
		DateOfResult:        test.DateOfResult,
		SampleIdentifier:    test.SampleIdentifier,
		SampleType:          test.SampleType,
		TestType:            testType,
		Result:              testResult,
		Status:              status,
		OutbreakID:          test.OutbreakID,
		TestedFor:           test.TestedFor,
		CreatedAt:           test.CreatedAt,
		CreatedBy:           test.CreatedBy,
		UpdatedAt:           test.UpdatedAt,
		UpdatedBy:           test.UpdatedBy,
		Person: models.LabTestCase{
			ID:             person.ID,
			FirstName:      person.FirstName,
			LastName:       person.LastName,
			MiddleName:     person.MiddleName,
			Gender:         person.Gender,
			Dob:            person.Dob,
			VisualID:       person.VisualID,
			Bhis:           person.Bhis,
			ReportingDate:  person.ReportingDate,
			CreatedAt:      person.CreatedAt,
			CreatedBy:      person.CreatedBy,
			Occupation:     person.Occupation,
			Classification: person.Classification,
			DateBecameCase: person.DateBecameCase,
			DateOfOnset:    person.DateOfOnset,
			RiskLevel:      person.RiskLevel,
			RiskReason:     person.RiskReason,
			Outcome:        person.Outcome,
			DateOfOutcome:  person.DateOfOutcome,
			Addresses:      person.Addresses,
			Documents:      person.Documents,
		},
		LabFacility: labFacility,
	}
	return labTest
}

// LabTestById searches for a lab test that has the specified id.
func (s *Store) LabTestById(ctx context.Context, id string) (models.LabTest, error) {
	labCol := s.Client.Database(s.Database).Collection(labCollection)
	filter := bson.D{{"_id", id}} //nolint:govet
	var rawLabTest models.RawLabTest
	err := labCol.FindOne(ctx, filter).Decode(&rawLabTest)
	if err != nil {
		return models.LabTest{}, MongoQueryErr{
			Reason: "LabTestById() error",
			Inner:  err,
		}
	}
	personCol := s.Client.Database(s.Database).Collection(personCollection)
	personFilter := bson.D{{"_id", rawLabTest.PersonId}} //nolint:govet
	var person models.Case
	personErr := personCol.FindOne(ctx, personFilter).Decode(&person)
	if personErr != nil {
		return models.LabTest{}, MongoQueryErr{
			Reason: "LabTestById() error querying person",
			Inner:  personErr,
		}
	}
	labTest := RawLabTestToLabTest(rawLabTest, person)
	return labTest, nil
}

func findCaseInCases(caseId string, cases []models.Case) models.Case {
	var person models.Case
	for i, _ := range cases { //nolint:gofmt,gosimple
		if caseId == cases[i].ID {
			person = cases[i]
			break
		}
	}
	return person
}

// FindLabTestsByDateRange returns all lab tests that occurred between two dates
func (s *Store) FindLabTestsByDateRange(ctx context.Context, startDate, endDate *time.Time, outbreakIDs []string) ([]models.LabTestReport, error) {
	var rawLabTests []models.RawLabTest
	var labTests []models.LabTestReport
	if startDate.After(*endDate) {
		return labTests, fmt.Errorf("startDate must be before endDate")
	}
	filter := bson.M{
		"deleted": false,
		"$and": bson.A{
			bson.M{"createdAt": bson.M{"$gte": startDate}},
			bson.M{"createdAt": bson.M{"$lt": endDate.Add(time.Hour * 24)}},
			bson.M{"outbreakId": bson.M{"$in": outbreakIDs}},
		},
	}

	labCol := s.Client.Database(s.Database).Collection(labCollection)
	cursor, err := labCol.Find(ctx, filter)
	if err != nil {
		return labTests, MongoQueryErr{
			Reason: "failed to fetch lab results",
			Inner:  err,
		}
	}

	if err := cursor.All(ctx, &rawLabTests); err != nil {
		return labTests, MongoQueryErr{
			Reason: fmt.Sprintf("error serializing lab tests from the db for ranges: %v, %v", startDate, endDate),
			Inner:  err,
		}
	}

	var personIds []string
	for i := range rawLabTests {
		personIds = append(personIds, rawLabTests[i].PersonId)
	}

	cases, err := s.FindCasesByPersonIds(ctx, personIds)
	if err != nil {
		return labTests, MongoQueryErr{
			Reason: "failed to fetch persons for lab tests",
			Inner:  err,
		}
	}

	for i := range rawLabTests {
		personID := rawLabTests[i].PersonId
		person := cases[personID]
		if person != nil {
			labTests = append(labTests, RawLabTestToLabReport(rawLabTests[i], *person))
		}
	}

	return labTests, nil
}
