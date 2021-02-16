package csv

import (
	"bz.moh.epi/godatatools/models"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

const layoutISO = "2006-01-02"

func WriteCases(writer *csv.Writer, cases []models.Case) error {
	headers := []string{
		"caseId",
		"bhisNumber",
		"dateOfReporting",
		"firstName",
		"lastName",
		"gender",
		"occupation",
		"ageInYears",
		"dateOfBirth",
		"classification",
		"dateBecameCase",
		"dateOfOnset",
		"riskLevel",
		"riskReason",
		"outcome",
		"pregnancyStatus",
		"dateOfOutcome",
		"address1_typeId",
		"address1_country",
		"address1_city",
		"address1_line1",
		"address1_line2",
		"address1_date",
		"address1_phoneNumber",
		"address2_typeId",
		"address2_country",
		"address2_city",
		"address2_line1",
		"address2_line2",
		"address2_date",
		"address2_phoneNumber",
		"address3_typeId",
		"address3_country",
		"address3_city",
		"address3_line1",
		"address3_line2",
		"address3_date",
		"address3_phoneNumber",
		"country_residence",
		"shows_symptoms",
		"fever",
		"soreThroat",
		"runnyNose",
		"cough",
		"vomiting",
		"nausea",
		"diarrhea",
		"shortnessOfBreath",
		"difficultyBreathing",
		"symptomsChills",
		"headache",
		"malaise",
		"anosmia",
		"aguesia",
		"bleeding",
		"jointMusclePain",
		"eyeFacialPain",
		"generalizedRash",
		"blurredVision",
		"abdominalPain",
		"caseType",
		"priorXDayExposureTravelledInternationally",
		"priorXDayExposureContactWithCase",
		"priorXDayExposureContactWithCaseDate",
		"priorXDayExposureInternationalDateTravelFrom",
		"priorXDayExposureInternationalDateTravelTo",
		"priorXDayExposureInternationaltravelcountries",
		"priorXDayExposureInternationalTravelCities",
		"typeOfTraveller",
		"purposeOfTravel",
		"flightNumber",
		"pcrTestinPast72Hours",
		"deathnCovContribution",
		"postMortemPerformed",
		"causeOfDeath",
		"respiratorySampleCollected",
		"respiratorySampleCollectedDate",
		"mechanicalVentilation",
		"hospitalization1_typeId",
		"hospitalization1_startDate",
		"hospitalization1_endDate",
		"hospitalization1_centerName",
		"hospitalization1_locationId",
		"hospitalization1_comments",
		"hospitalization2_typeId",
		"hospitalization2_startDate",
		"hospitalization2_endDate",
		"hospitalization2_centerName",
		"hospitalization2_locationId",
		"hospitalization2_comments",
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error: writing header for csv: %w", err)
	}

	for _, c := range cases {
		var record []string
		record = append(record, c.VisualID)
		record = append(record, fmt.Sprintf("%v", c.Bhis))
		record = append(record, c.ReportingDate.Format(layoutISO))
		record = append(record, c.FirstName)
		record = append(record, c.LastName)
		record = append(record, c.Gender)
		record = append(record, c.Occupation)
		record = append(record, fmt.Sprintf("%v", c.Age.Years))
		record = append(record, c.Dob.Format(layoutISO))
		record = append(record, c.Classification)
		var dateBecameCase string
		if c.DateBecameCase != nil {
			dateBecameCase = c.DateBecameCase.Format(layoutISO)
		}
		record = append(record, dateBecameCase)
		var dateOnset string
		if c.DateOfOnset != nil {
			dateOnset = c.DateOfOnset.Format(layoutISO)
		}
		record = append(record, dateOnset)
		record = append(record, c.RiskLevel)
		record = append(record, c.RiskReason)
		record = append(record, c.Outcome)
		record = append(record, c.PregnancyStatus)
		var dateOutcome string
		if c.DateOfOutcome != nil {
			dateOutcome = c.DateOfOutcome.Format(layoutISO)
		}
		record = append(record, dateOutcome)
		for _, a := range c.Addresses {
			record = append(record, a.TypeId)
			record = append(record, a.Country)
			record = append(record, a.City)
			record = append(record, a.AddressLine1)
			record = append(record, a.AddressLine2)
			var addressDate string
			if a.Date != nil {
				addressDate = a.Date.Format(layoutISO)
			}
			record = append(record, addressDate)
			record = append(record, a.PhoneNumber)
		}

		// Questionnaire
		questionnaire := c.Questionnaire
		var residence string
		if len(questionnaire.CountryResidence) > 0 {
			residence = questionnaire.CountryResidence[0].Value
		}
		record = append(record, residence)
		var showsSymptoms string
		if len(questionnaire.ShowsSymptoms) > 0 {
			showsSymptoms = questionnaire.ShowsSymptoms[0].Value
		}
		record = append(record, showsSymptoms)
		var fever string
		if len(questionnaire.Fever) > 0 {
			fever = questionnaire.Fever[0].Value
		}
		record = append(record, fever)
		var soreThroat string
		if len(questionnaire.SoreThroat) > 0 {
			soreThroat = questionnaire.SoreThroat[0].Value
		}
		record = append(record, soreThroat)
		var runnyNose string
		if len(questionnaire.RunnyNose) > 0 {
			runnyNose = questionnaire.RunnyNose[0].Value
		}
		record = append(record, runnyNose)
		var cough string
		if len(questionnaire.Cough) > 0 {
			cough = questionnaire.Cough[0].Value
		}
		record = append(record, cough)
		var vomiting string
		if len(questionnaire.Vomiting) > 0 {
			vomiting = questionnaire.Vomiting[0].Value
		}
		record = append(record, vomiting)
		var nausea string
		if len(questionnaire.Nausea) > 0 {
			nausea = questionnaire.Nausea[0].Value
		}
		record = append(record, nausea)
		var diarrhea string
		if len(questionnaire.Diarrhea)  > 0 {
			diarrhea = questionnaire.Diarrhea[0].Value
		}
		record = append(record, diarrhea)
		var shortnessBreath string
		if len(questionnaire.ShortnessOfBreath) > 0 {
			shortnessBreath = questionnaire.ShortnessOfBreath[0].Value
		}
		record = append(record, shortnessBreath)
		var difficultyBreathing string
		if len(questionnaire.DifficultyBreathing) > 0 {
			difficultyBreathing = questionnaire.DifficultyBreathing[0].Value
		}
		record = append(record, difficultyBreathing)
		var chills string
		if len(questionnaire.SymptomsChills) > 0 {
			chills = questionnaire.SymptomsChills[0].Value
		}
		record = append(record, chills)
		var headache string
		if len(questionnaire.Headache) > 0 {
			headache = questionnaire.Headache[0].Value
		}
		record = append(record, headache)
		var malaise string
		if len(questionnaire.Malaise) > 0 {
			malaise = questionnaire.Malaise[0].Value
		}
		record = append(record, malaise)
		var anosmia string
		if len(questionnaire.Anosmia) > 0 {
			anosmia = questionnaire.Anosmia[0].Value
		}
		record = append(record, anosmia)
		var aguesia string
		if len(questionnaire.Aguesia) > 0 {
			aguesia = questionnaire.Aguesia[0].Value
		}
		record = append(record, aguesia)
		var bleeding string
		if len(questionnaire.Bleeding) > 0 {
			bleeding = questionnaire.Bleeding[0].Value
		}
		record = append(record, bleeding)
		var musclePain string
		if len(questionnaire.JointMusclePain) > 0 {
			musclePain = questionnaire.JointMusclePain[0].Value
		}
		record = append(record, musclePain)
		var eyeFacialPain string
		if len(questionnaire.EyeFacialPain) > 0 {
			eyeFacialPain = questionnaire.EyeFacialPain[0].Value
		}
		record = append(record, eyeFacialPain)
		var generalizedRash string
		if len(questionnaire.GeneralizedRash) > 0 {
			generalizedRash = questionnaire.GeneralizedRash[0].Value
		}
		record = append(record, generalizedRash)
		var blurredVision string
		if len(questionnaire.BlurredVision) > 0 {
			blurredVision = questionnaire.BlurredVision[0].Value
		}
		record = append(record, blurredVision)
		var abdominalPain string
		if len(questionnaire.AbdominalPain) > 0 {
			abdominalPain = questionnaire.AbdominalPain[0].Value
		}
		record = append(record, abdominalPain)
		record = append(record, questionnaire.CaseType)
		var internationalTravel string
		if len(questionnaire.PriorXdayExposureTravelledInternationally) > 0 {
			internationalTravel = questionnaire.PriorXdayExposureTravelledInternationally[0].Value
		}
		record = append(record,internationalTravel)
		var contactWithCase string
		if len(questionnaire.PriorXdayExposureContactWithCase) > 0 {
			contactWithCase = questionnaire.PriorXdayExposureContactWithCase[0].Value
		}
		record = append(record, contactWithCase)
		var contactDate string
		if len(questionnaire.PriorXdayExposureContactWithCaseDate) > 0 && questionnaire.PriorXdayExposureContactWithCaseDate[0].Value != nil {
			rawContactDate := questionnaire.PriorXdayExposureContactWithCaseDate[0].Value
			typeOfContactDate := reflect.TypeOf(rawContactDate)
			if typeOfContactDate.String() == "string" {
				contactDate = questionnaire.PriorXdayExposureContactWithCaseDate[0].Value.(string)
			} else {
				date := questionnaire.PriorXdayExposureContactWithCaseDate[0].Value.(primitive.DateTime)
				contactDate = date.Time().Format(layoutISO)
			}

		}
		record = append(record, contactDate)
		var dateTravelFrom string
		if len(questionnaire.PriorXdayExposureInternationalDateTravelFrom) > 0 && questionnaire.PriorXdayExposureInternationalDateTravelFrom[0].Value != nil {
			dateTravelFrom = questionnaire.PriorXdayExposureInternationalDateTravelFrom[0].Value.(string)
		}
		record = append(record, dateTravelFrom)
		var intlTravelDate string
		if len(questionnaire.PriorXdayExposureInternationalDatetravelTo) > 0 {
			intlTravelDate = questionnaire.PriorXdayExposureInternationalDatetravelTo[0].Value
		}
		record = append(record,intlTravelDate)
		var travelCountries string
		if len(questionnaire.PriorXdayExposureInternationaltravelcountries) > 0 {
			travelCountries = questionnaire.PriorXdayExposureInternationaltravelcountries[0].Value
		}
		record = append(record, travelCountries)
		var travelCities string
		if len(questionnaire.PriorXdayExposureInternationalTravelCities) > 0 {
			travelCities = questionnaire.PriorXdayExposureInternationalTravelCities[0].Value
		}
		record = append(record, travelCities)
		var typeOfTraveller string
		if len(questionnaire.TypeOfTraveller)> 0 {
			typeOfTraveller = questionnaire.TypeOfTraveller[0].Value
		}
		record = append(record, typeOfTraveller)
		var purposeOfTravel string
		if len(questionnaire.PurposeOfTravel) > 0 {
			purposeOfTravel = questionnaire.PurposeOfTravel[0].Value
		}
		record = append(record, purposeOfTravel)
		var flightNumber string
		if len(questionnaire.FlightNumber) > 0 {
			flightNumber = questionnaire.FlightNumber[0].Value
		}
		record = append(record, flightNumber)
		var pcrTest string
		if len(questionnaire.PcrTestInPast72Hours) > 0 {
			pcrTest = questionnaire.PcrTestInPast72Hours[0].Value
		}
		record = append(record, pcrTest)
		var deathContrib string
		if len(questionnaire.DeathContrib) > 0 {
			deathContrib = questionnaire.DeathContrib[0].Value
		}
		record = append(record, deathContrib)
		var postMortem string
		if len(questionnaire.PostMortem) > 0 {
			postMortem = questionnaire.PostMortem[0].Value
		}
		record = append(record, postMortem)
		var causeOfDeath string
		if len(questionnaire.CauseOfDeath) > 0 {
			causeOfDeath = questionnaire.CauseOfDeath[0].Value
		}
		record = append(record, causeOfDeath)
		var respSampleCollected string
		if len(questionnaire.RespSampleCollected) > 0 {
			respSampleCollected = questionnaire.RespSampleCollected[0].Value
		}
		record = append(record, respSampleCollected)
		var respSampleDateCollected string
		if len(questionnaire.RespiratorySampleDateCollected) > 0 && questionnaire.RespiratorySampleDateCollected[0].Value != nil {
			rawSampleDate := questionnaire.RespiratorySampleDateCollected[0].Value
			typeOfSampleDate := reflect.TypeOf(rawSampleDate)
			if typeOfSampleDate.String() == "string" {
				contactDate = questionnaire.RespiratorySampleDateCollected[0].Value.(string)
			} else {
				sampleDate := questionnaire.RespiratorySampleDateCollected[0].Value.(primitive.DateTime)
				respSampleDateCollected = sampleDate.Time().Format(layoutISO)
			}
		}
		record = append(record, respSampleDateCollected)
		var mechanicalVentilation string
		if len(questionnaire.MechanicalVentilation) > 0 {
			mechanicalVentilation = questionnaire.MechanicalVentilation[0].Value
		}
		record = append(record, mechanicalVentilation)

		// Process the Hospitalizations data
		var hospitalizations []models.Hospitalization
		if c.Hospitalizations != nil {
			hospitalizations = c.Hospitalizations
		}
		for _, h := range hospitalizations {
			record = append(record, h.TypeId)
			var startDate string
			if h.StartDate != nil {
				startDate = h.StartDate.Format(layoutISO)
			}
			record = append(record, startDate)
			var endDate string
			if h.EndDate != nil {
				endDate = h.EndDate.Format(layoutISO)
			}
			record = append(record, endDate)
			record = append(record, h.CenterName)
			record = append(record, h.LocationId)
			record = append(record, h.Comments)
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write csv record(%s): %w", c.VisualID, err)
		}
	}
	return nil
}
