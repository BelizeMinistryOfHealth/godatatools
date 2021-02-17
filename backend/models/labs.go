package models

import "time"

// RawLabTest is the representation of how the lab test is stored in Mongo.
type RawLabTest struct {
	ID                  string    `bson:"_id" json:"id"`
	PersonType          string    `json:"personType"`
	PersonId            string    `json:"personId"`
	DateSampleTaken     time.Time `bson:"dateSampleTaken" json:"dateSampleTaken"`
	DateSampleDelivered time.Time `bson:"dateSampleDelivered" json:"dateSampleDelivered"`
	DateTesting         time.Time `bson:"dateTesting" json:"dateTesting"`
	DateOfResult        time.Time `bson:"dateOfResult" json:"dateOfResult"`
	SampleIdentifier    string    `json:"sampleIdentifier"`
	SampleType          string    `json:"sampleType"`
	TestType            string    `json:"testType"`
	Result              string    `json:"result"`
	Status              string    `json:"status"`
	OutbreakID          string    `json:"outbreakId"`
	TestedFor           string    `json:"testedFor"`
	CreatedAt           time.Time `bson:"createdAt" json:"createdAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedAt           time.Time `bson:"updatedAt" json:"updatedAt"`
}

type LabTest struct {
	ID                  string    `json:"id"`
	LabName             string    `json:"string"`
	PersonType          string    `json:"personType"`
	DateSampleTaken     time.Time `json:"dateSampleTaken"`
	DateSampleDelivered time.Time `json:"dateSampleDelivered"`
	DateTesting         time.Time `json:"dateTesting"`
	DateOfResult        time.Time `json:"dateOfResult"`
	SampleIdentifier    string    `json:"sampleIdentifier"`
	SampleType          string    `json:"sampleType"`
	TestType            string    `json:"testType"`
	Result              string    `json:"result"`
	Status              string    `json:"status"`
	OutbreakID          string    `json:"outbreakId"`
	TestedFor           string    `json:"testedFor"`
	CreatedAt           time.Time `json:"createdAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedAt           time.Time `json:"updatedAt"`
	Person              Case      `json:"person"`
}
