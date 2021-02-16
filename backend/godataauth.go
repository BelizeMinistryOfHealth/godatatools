package godatatools

import (
	"encoding/json"
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
		username string
		password string
	}

	var creds authCreds

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "could not decode body", http.StatusBadRequest)
	}
	token, err := s.GoData.GetToken(creds.username, creds.password)
	if err != nil {
		http.Error(w, "failed to authenticate", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
