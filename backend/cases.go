package godatatools

import (
	"bytes"
	csv2 "bz.moh.epi/godatatools/csv"
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
