package csv

import (
	"bz.moh.epi/godatatools/models"
	"encoding/csv"
	"fmt"
)

func WriteLabs(writer *csv.Writer, labs []models.LabTestReport) error {
	headers := []string{
		"ID",
		"createdAt",
		"createdBy",
		"updatedOn",
		"updatedBy",
		"personId",
		"dateOfBirth",
		"dateSampleTaken",
		"dateSampleDelivered",
		"dateSampleTested",
		"dateOfResult",
		"labName",
		"sampleLabId",
		"sampleType",
		"testType",
		"testedFor",
		"result",
		"status",
		"caseId",
		"personType",
		"lastName",
		"firstName",
		"middleName",
		"dateOfOnset",
		"dateReported",
		"addressType",
		"country",
		"city",
		"district",
		"addressLine1",
		"postalCode",
		"ssn",
		"bhis",
		"passport",
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error: writing header for csv: %w", err)
	}

	for i := range labs {
		var record []string
		lab := labs[i]
		record = append(record, lab.ID)
		record = append(record, lab.CreatedAt.Format(layoutISO))
		record = append(record, lab.CreatedBy)
		if lab.UpdatedAt == nil {
			record = append(record, "")
		} else {
			record = append(record, lab.UpdatedAt.Format(layoutISO))
		}
		record = append(record, lab.UpdatedBy)
		record = append(record, lab.Person.ID)
		record = append(record, lab.Person.Dob.Format(layoutISO))
		record = append(record, lab.DateSampleTaken.Format(layoutISO))
		if lab.DateSampleDelivered == nil {
			record = append(record, "")
		} else {
			record = append(record, lab.DateSampleDelivered.Format(layoutISO))
		}
		if lab.DateTesting == nil {
			record = append(record, "")
		} else {
			record = append(record, lab.DateTesting.Format(layoutISO))
		}
		if lab.DateOfResult == nil {
			record = append(record, "")
		} else {
			record = append(record, lab.DateOfResult.Format(layoutISO))
		}
		record = append(record, lab.LabName)
		record = append(record, lab.SampleIdentifier)
		record = append(record, lab.SampleType)
		record = append(record, lab.TestType)
		record = append(record, lab.TestedFor)
		record = append(record, lab.Result)
		record = append(record, lab.Status)
		record = append(record, lab.Person.VisualID)
		record = append(record, lab.PersonType)
		record = append(record, lab.Person.LastName)
		record = append(record, lab.Person.FirstName)
		record = append(record, lab.Person.MiddleName)
		if lab.Person.DateOfOnset == nil {
			record = append(record, "")
		} else {
			record = append(record, lab.Person.DateOfOnset.Format(layoutISO))
		}
		record = append(record, lab.Person.ReportingDate.Format(layoutISO))
		if lab.Person.Addresses == nil {
			record = append(record, "")
			record = append(record, "")
			record = append(record, "")
			record = append(record, "")
			record = append(record, "")
			record = append(record, "")
		} else {
			address := lab.Person.Addresses[0]
			record = append(record, address.TypeId)
			record = append(record, address.Country)
			record = append(record, address.City)
			record = append(record, "")
			record = append(record, address.AddressLine1)
			record = append(record, address.PostalCode)
		}

		if lab.Person.Documents == nil {
			record = append(record, "")
			record = append(record, "")
			record = append(record, "")
		} else {
			var ssn = ""
			var bhis = ""
			var passport = ""
			for i := range lab.Person.Documents {
				if lab.Person.Documents[i].Type == models.SSN {
					ssn = lab.Person.Documents[i].Number
				}
				if lab.Person.Documents[i].Type == models.BHIS {
					bhis = lab.Person.Documents[i].Number
				}
				if lab.Person.Documents[i].Type == models.PASSPORT {
					passport = lab.Person.Documents[i].Number
				}
			}
			record = append(record, fmt.Sprintf("%09s", ssn))
			record = append(record, bhis)
			record = append(record, passport)
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write csv record (%v), ", record)
		}
	}

	return nil
}
