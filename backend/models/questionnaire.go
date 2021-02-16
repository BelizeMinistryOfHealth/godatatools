package models

import "time"

type QuestionnaireAnswer struct {
	Value string `json:"value"`
}

type DateAnswer struct {
	Value interface{} `json:"value"`
}

type CaseForm struct {
	Value []string `json:"value"`
}

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Address struct {
	TypeId       string       `json:"typeId"`
	Country      string       `json:"country"`
	City         string       `json:"city"`
	AddressLine1 string       `json:"addressLine1"`
	AddressLine2 string       `json:"addressLine2"`
	Date         *time.Time   `bson:"date" json:"date"`
	PhoneNumber  string       `json:"phoneNumber"`
	LocationId   string       `json:"locationId"`
	GeoLocation  *GeoLocation `json:"geoLocation,omitempty"`
}

type PersonAge struct {
	Years  int `json:"years"`
	Months int `json:"months"`
}

// GoDataQuestionnaire represents the GoData questionnaire. GoData stores these as a flat list.
// The CaseForm identifies the forms, and GoData uses this to extract the fields for each form from the
// flat list of questions.
type Questionnaire struct {
	CaseForm                                      []CaseForm            `bson:"Case_WhichForm" json:"Case_WhichForm"`
	DataCollectorName                             []QuestionnaireAnswer `bson:"FA0_datacollector_name" json:"FA0_datacollector_name"`
	CountryResidence                              []QuestionnaireAnswer `bson:"FA0_case_countryresidence" json:"FA0_case_countryresidence"`
	ShowsSymptoms                                 []QuestionnaireAnswer `bson:"FA0_symptoms_caseshowssymptoms" json:"FA0_symptoms_caseshowssymptoms"`
	Fever                                         []QuestionnaireAnswer `bson:"FA0_symptom_fever" json:"FA0_symptom_fever"`
	SoreThroat                                    []QuestionnaireAnswer `bson:"FA0_symptom_sorethroat" json:"FA0_symptom_sorethroat"`
	RunnyNose                                     []QuestionnaireAnswer `bson:"FA0_symptom_runnynose" json:"FA0_symptom_runnynose"`
	Cough                                         []QuestionnaireAnswer `bson:"FA0_symptom_cough" json:"FA0_symptom_cough"`
	Vomiting                                      []QuestionnaireAnswer `bson:"FA0_symptom_vomiting" json:"FA0_symptom_vomiting"`
	Nausea                                        []QuestionnaireAnswer `bson:"FA0_symptom_nausea" json:"FA0_symptom_nausea"`
	Diarrhea                                      []QuestionnaireAnswer `bson:"FA0_symptom_diarrhea" json:"FA0_symptom_diarrhea"`
	ShortnessOfBreath                             []QuestionnaireAnswer `bson:"FA0_symptom_shortnessofbreath" json:"FA0_symptom_shortnessofbreath"`
	DifficultyBreathing                           []QuestionnaireAnswer `bson:"FA0_symptom_difficulty_breathing" json:"FA0_symptom_difficulty_breathing"`
	SymptomsChills                                []QuestionnaireAnswer `bson:"FA0_symptom_chills" json:"FA0_symptom_chills"`
	Headache                                      []QuestionnaireAnswer `bson:"FA0_symptom_headache" json:"FA0_symptom_headache"`
	Malaise                                       []QuestionnaireAnswer `bson:"FA0_symptom_malaise" json:"FA0_symptom_malaise"`
	Anosmia                                       []QuestionnaireAnswer `bson:"FA0_symptom_anosmia" json:"FA0_symptom_anosmia"`
	Aguesia                                       []QuestionnaireAnswer `bson:"FA0_symptom_aguesia" json:"FA0_symptom_aguesia"`
	Bleeding                                      []QuestionnaireAnswer `bson:"FA0_symptom_bleeding" json:"FA0_symptom_bleeding"`
	JointMusclePain                               []QuestionnaireAnswer `bson:"FA0_symptom_joint_muscle_pain" json:"FA0_symptom_joint_muscle_pain"`
	EyeFacialPain                                 []QuestionnaireAnswer `bson:"FA0_symptom_eye_facial_pain" json:"FA0_symptom_eye_facial_pain"`
	GeneralizedRash                               []QuestionnaireAnswer `bson:"FA0_symptom_generalized_rash" json:"FA0_symptom_generalized_rash"`
	BlurredVision                                 []QuestionnaireAnswer `bson:"FA0_symptom_blurred_vision" json:"FA0_symptom_blurred_vision"`
	AbdominalPain                                 []QuestionnaireAnswer `bson:"FA0_symptom_abdominal_pain" json:"FA0_symptom_abdominal_pain"`
	CaseType                                      string                `bson:"case_type" json:"case_type"`
	Ssn                                           []QuestionnaireAnswer `bson:"FA0_caseidentifier_socialnumber" json:"FA0_caseidentifier_socialnumber"`
	PriorXdayExposureTravelledInternationally     []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_travelledinternationally" json:"FA0_priorXdayexposure_travelledinternationally"`
	PriorXdayExposureContactWithCase              []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_contactwithcase" json:"FA0_priorXdayexposure_contactwithcase"`
	PriorXdayExposureContactWithCaseDate          []DateAnswer          `bson:"FA0_priorXdayexposure_contactwithcasedate" json:"FA0_priorXdayexposure_contactwithcasedate"`
	PriorXdayExposureInternationalDateTravelFrom  []DateAnswer          `bson:"FA0_priorXdayexposure_internationaldatetravelfrom" json:"FA0_priorXdayexposure_internationaldatetravelfrom"`
	PriorXdayExposureInternationalDatetravelTo    []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_internationaldatetravelto" json:"FA0_priorXdayexposure_internationaldatetravelto"`
	PriorXdayExposureInternationaltravelcountries []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_internationaltravelcountries" json:"FA0_priorXdayexposure_internationaltravelcountries"`
	PriorXdayExposureInternationalTravelCities    []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_internationaltravelcities" json:"FA0_priorXdayexposure_internationaltravelcities"`
	TypeOfTraveller                               []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_typeoftraveler" json:"FA0_priorXdayexposure_typeoftraveler"`
	PurposeOfTravel                               []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_purposeoftravel" json:"FA0_priorXdayexposure_purposeoftravel"`
	FlightNumber                                  []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_flightnumber" json:"FA0_priorXdayexposure_flightnumber"`
	PcrTestInPast72Hours                          []QuestionnaireAnswer `bson:"FA0_priorXdayexposure_tookpcrtest_past72hours" json:"FA0_priorXdayexposure_tookpcrtest_past72hours"`
	DeathContrib                                  []QuestionnaireAnswer `bson:"FA2_outcome_deathnCoVcontribution" json:"FA2_outcome_deathnCoVcontribution"`
	PostMortem                                    []QuestionnaireAnswer `bson:"FA2_outcome_postmortemperformed" json:"FA2_outcome_postmortemperformed"`
	CauseOfDeath                                  []QuestionnaireAnswer `bson:"FA2_symptoms_causeofdeath" json:"FA2_symptoms_causeofdeath"`
	RespSampleCollected                           []QuestionnaireAnswer `bson:"FA0_respiratorysample_collectedYN" json:"FA0_respiratorysample_collectedYN"`
	RespiratorySampleDateCollected                []DateAnswer          `bson:"FA0_respiratorysample_datecollected" json:"FA0_respiratorysample_datecollected"`
	MechanicalVentilation                         []QuestionnaireAnswer `bson:"FA0_clinicalcomplications_mechanicalventilation" json:"FA0_clinicalcomplications_mechanicalventilation"`
}
