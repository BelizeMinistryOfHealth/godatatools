package godatatools

import (
	"bytes"
	csv2 "bz.moh.epi/godatatools/csv"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const isoLayout string = "2006-01-02"

func (s Server) LabTestsByDateRange(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "GET only", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	startDateQuery := query.Get("startDate")
	if len(startDateQuery) == 0 {
		http.Error(w, "startDate was not provided", http.StatusBadRequest)
		return
	}
	startDate, err := time.Parse(isoLayout, startDateQuery)
	if err != nil {
		log.WithError(err).Errorf("error parsing startDate: %s", startDateQuery)
		http.Error(w, "invalid startDate", http.StatusBadRequest)
		return
	}
	endDateQuery := query.Get("endDate")
	if len(endDateQuery) == 0 {
		log.WithError(err).Errorf("error parsing endDate: %s", endDateQuery)
		http.Error(w, "endDate was not provided", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse(isoLayout, endDateQuery)
	if err != nil {
		http.Error(w, "invalid endDate", http.StatusBadRequest)
		return
	}

	if startDate.After(endDate) {
		http.Error(w, "startDate must be before endDate", http.StatusBadRequest)
		return
	}

	labTests, err := s.DbRepository.FindLabTestsByDateRange(r.Context(), &startDate, &endDate)
	if err != nil {
		log.WithError(err).Errorf("failed to retrieve lab tests")
		http.Error(w, "failed to retrieve lab tests", http.StatusInternalServerError)
		return
	}

	b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(b)
	if err := csv2.WriteLabs(csvWriter, labTests); err != nil {
		log.WithError(err).Error("failed to convert labtests to csv")
		http.Error(w, "error converting lab tests to csv", http.StatusInternalServerError)
		return
	}
	csvWriter.Flush()
	if _, err := w.Write(b.Bytes()); err != nil {
		log.WithError(err).Error("failed to stream lab tests csv file")
		http.Error(w, "error streaming lab tests csv file", http.StatusInternalServerError)
		return
	}
	return
}
