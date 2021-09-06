package godatatools

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s Server) AuthWithGodata(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	type authCreds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var creds authCreds

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
	}
	token, err := s.GoData.GetToken(creds.Username, creds.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"godata": s.GoData,
			"creds":  creds,
		}).WithError(err).Error("authentication failed")
		http.Error(w, "failed to authenticate", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
