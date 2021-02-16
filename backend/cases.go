package godatatools

import (
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
	json.NewEncoder(w).Encode(cases)

}
