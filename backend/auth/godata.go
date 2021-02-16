package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GoData struct {
	BaseURL string
}

type goDataAuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (g GoData) GetToken(username, password string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{"username": username, "password": password})
	if err != nil {
		return "", err
	}
	req, err := http.Post(fmt.Sprintf("%s/api/oauth/token", g.BaseURL), "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	var authResp *goDataAuthResponse
	if err := json.NewDecoder(req.Body).Decode(&authResp); err != nil {
		return "", err
	}
	if req.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth with godata failed: status: %d", req.StatusCode)
	}
	return authResp.AccessToken, nil
}
