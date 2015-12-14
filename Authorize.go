package b2

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

var authURL string = "https://api.backblaze.com/b2api/v1/b2_authorize_account"

type B2 struct {
	AccountID          string `json:"accountId"`
	ApiURL             string `json:"apiUrl"`
	AuthorizationToken string `json:"authorizationToken"`
	DownloadURL        string `json:"downloadUrl"`

	client *http.Client
}

func Authorize(acountID, secretKey string) (*B2, error) {
	id_and_key := []byte(acountID + ":" + secretKey)
	basic_auth_string := "Basic " + base64.StdEncoding.EncodeToString(id_and_key)

	client := &http.Client{}
	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", basic_auth_string)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var b2 B2
	for dec.More() {
		if err := dec.Decode(&b2); err != nil {
			return nil, err
		}
	}

	b2.client = client
	return &b2, nil
}
