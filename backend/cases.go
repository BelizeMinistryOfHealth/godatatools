package godatatools

import (
	"bytes"
	csv2 "bz.moh.epi/godatatools/csv"
	"bz.moh.epi/godatatools/models"
	"encoding/csv"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s Server) CasesByOutbreak(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	outbreakId := query.Get("outbreakId")
	if len(outbreakId) == 0 {
		http.Error(w, "outbreakId was not provided", http.StatusBadRequest)
		return
	}

	cases, err := s.DbRepository.FindCasesByOutbreak(r.Context(), outbreakId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"outbreakId": outbreakId,
		}).WithError(err).Error("could not retrieve cases for the outbreak")
		return
	}
	b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(b)
	if err := csv2.WriteCases(csvWriter, cases); err != nil {
		log.WithError(err).Error("failed to convert cases to csv")
		http.Error(w, "error converting cases to csv", http.StatusInternalServerError)
		return
	}
	csvWriter.Flush()
	if _, err := w.Write(b.Bytes()); err != nil {
		log.WithError(err).Error("failed to stream file")
		http.Error(w, "error streaming file", http.StatusInternalServerError)
		return
	}
	return
}

func (s Server) AllOutbreaks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusBadRequest)
		return
	}

	outbreaks, err := s.DbRepository.ListOutbreaks(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.WithError(err).Error("could not retrieve the outbreaks")
		return
	}
	json.NewEncoder(w).Encode(outbreaks)
}

func (s Server) LabTestResults(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusBadRequest)
	}

	// Get the case id from the url path
	query := r.URL.Query()
	firstName := query.Get("firstName")
	if len(firstName) == 0 {
		http.Error(w, "a first name must be provided", http.StatusBadRequest)
		return
	}
	lastName := query.Get("lastName")
	if len(lastName) == 0 {
		http.Error(w, "a last name must be provided", http.StatusBadRequest)
		return
	}

	labTests, err := s.DbRepository.LabTestsByCaseName(r.Context(), firstName, lastName)
	if err != nil {
		log.WithFields(log.Fields{
			"firstName": firstName,
			"lastName":  lastName,
		}).WithError(err).Error("error retrieving lab tests by first and last names")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if labTests == nil {
		labTests = []models.LabTest{}
	}

	if err := json.NewEncoder(w).Encode(labTests); err != nil {
		log.WithFields(log.Fields{
			"firstName": firstName,
			"lastName":  lastName,
		}).WithError(err).Error("error encoding lab tests results")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
