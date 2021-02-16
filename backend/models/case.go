package models

import "time"

type Hospitalization struct {
	TypeId     string     `json:"typeId"`
	StartDate  *time.Time `bson:"startDate" json:"startDate"`
	EndDate    *time.Time `bson:"endDate" json:"endDate,omitempty"`
	CenterName string     `json:"centerName"`
	LocationId string     `json:"locationId"`
	Comments   string     `json:"comments"`
}

type Case struct {
	ID               string            `bson:"_id" json:"id"`
	VisualID         string            `json:"visualId"`
	Bhis             int               `json:"bhis"`
	ReportingDate    time.Time         `bson:"dateOfReporting" json:"dateOfReporting"`
	CreatedAt        time.Time         `json:"createdAt"`
	CreatedBy        string            `json:"createdBy"`
	FirstName        string            `json:"firstName"`
	LastName         string            `json:"lastName"`
	Gender           string            `json:"gender"`
	Occupation       string            `json:"occupation"`
	Age              PersonAge         `json:"age"`
	Dob              time.Time         `json:"dob"`
	Classification   string            `json:"classification"`
	DateBecameCase   *time.Time        `json:"dateBecomeCase"`
	DateOfOnset      *time.Time        `json:"dateOfOnset"`
	RiskLevel        string            `json:"riskLevel"`
	RiskReason       string            `json:"riskReason"`
	Outcome          string            `json:"outcomeId"`
	PregnancyStatus  string            `json:"pregnancyStatus"`
	DateOfOutcome    *time.Time        `json:"dateOfOutcome"`
	Addresses        []Address         `json:"addresses"`
	Questionnaire    Questionnaire     `bson:"questionnaireAnswers" json:"questionnaireAnswers"`
	Hospitalizations []Hospitalization `bson:"dateRanges" json:"dateRanges,omitempty"`
}

type Outbreak struct {
	ID   string `bson:"_id" json:"_id"`
	Name string `json:"name"`
}
