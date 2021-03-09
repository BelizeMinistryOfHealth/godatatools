package models

import "time"

// RawLabTest is the representation of how the lab test is stored in Mongo.
type RawLabTest struct {
	ID                  string     `bson:"_id" json:"id"`
	PersonType          string     `json:"personType"`
	PersonId            string     `json:"personId"`
	DateSampleTaken     *time.Time `bson:"dateSampleTaken" json:"dateSampleTaken"`
	DateSampleDelivered *time.Time `bson:"dateSampleDelivered" json:"dateSampleDelivered"`
	DateTesting         *time.Time `bson:"dateTesting" json:"dateTesting"`
	DateOfResult        *time.Time `bson:"dateOfResult" json:"dateOfResult"`
	SampleIdentifier    string     `json:"sampleIdentifier"`
	SampleType          string     `json:"sampleType"`
	TestType            string     `json:"testType"`
	Result              string     `json:"result"`
	Status              string     `json:"status"`
	OutbreakID          string     `json:"outbreakId"`
	TestedFor           string     `json:"testedFor"`
	CreatedAt           time.Time  `bson:"createdAt" json:"createdAt"`
	CreatedBy           string     `json:"createdBy"`
	UpdatedAt           time.Time  `bson:"updatedAt" json:"updatedAt"`
}

type Person struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	Dob       time.Time `json:"dob"`
	Age       int       `json:"age"`
}

type LabTest struct {
	ID                  string      `json:"id"`
	LabName             string      `json:"labName"`
	PersonType          string      `json:"personType"`
	DateSampleTaken     *time.Time  `json:"dateSampleTaken"`
	DateSampleDelivered *time.Time  `json:"dateSampleDelivered"`
	DateTesting         *time.Time  `json:"dateTesting"`
	DateOfResult        *time.Time  `json:"dateOfResult"`
	SampleIdentifier    string      `json:"sampleIdentifier"`
	SampleType          string      `json:"sampleType"`
	TestType            string      `json:"testType"`
	Result              string      `json:"result"`
	Status              string      `json:"status"`
	OutbreakID          string      `json:"outbreakId"`
	TestedFor           string      `json:"testedFor"`
	CreatedAt           time.Time   `json:"createdAt"`
	CreatedBy           string      `json:"createdBy"`
	UpdatedAt           time.Time   `json:"updatedAt"`
	Person              Person      `json:"person"`
	LabFacility         LabFacility `json:"labFacility"`
}

type LabFacility struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var LabFacilities = []LabFacility{
	{ID: "5fc2d66b-8af8-42eb-a47a-c56fdd42264a", Name: "PGIA"},
	{ID: "b11d5713-3456-4de7-b85a-e30d5f4566f2", Name: "2020 Outbreak - N/A - local"},
	{ID: "d54c3aa5-7f43-4733-a482-32d4f8d0b8c4", Name: "2020 Outbreak - N/A - Main"},
	{ID: "d6e93ee2-9e47-4f21-8c11-d4bc06041020", Name: "Belize Medical Associates"},
	{ID: "5b3618b3-6313-47cd-90c8-9020250eddef", Name: "Belize Healthcare Partners Limited"},
	{ID: "0088b606-f472-46f5-96b7-9cab8ac3a69e", Name: "Belize Medical Associates"},
	{ID: "9a8e1b83-237e-42c9-aeca-8f96a073aca2", Name: "TEST - N/A"},
	{ID: "5996a16e-b0e6-4d95-bec9-9559ad337014", Name: "Belize Healthcare Partners Limited"},
	{ID: "af1ed81d-07a1-4ff8-b218-bde2f99a746e", Name: "St. Luke Hospital"},
	{ID: "74e5f746-a7ee-4654-b55f-05388ae9fd47", Name: "Belmopan Medical Center"},
	{ID: "13774a30-842f-465a-9c05-50e57467be9b", Name: "TEST - N/A"},
	{ID: "34d42724-7967-43d1-8af8-28f7f65a2749", Name: "Belize Diagnostic Center"},
	{ID: "b28b0af9-4b25-4a61-bba3-fddea0d107ac", Name: "Northern Medical Specialty Plaza"},
	{ID: "2ddfddee-5fb7-4825-8d0d-3586545a0a16", Name: "Caring Hands"},
	{ID: "52b22cd9-5dc8-4e5b-be0c-e53dc45afbe7", Name: "Belize Specialists Hospital"},
	{ID: "fe40d930-db98-4e1f-8303-0ed38b6f7621", Name: "Archangel Medical Center"},
	{ID: "50ed1fc6-8812-4f30-862b-65b9699dac8a", Name: "Placencia Medical Services and Southern Clinical Lab"},
	{ID: "25ed7176-51b6-400b-9713-718dedc4262c", Name: "Genysis Medical Laboratory"},
}
