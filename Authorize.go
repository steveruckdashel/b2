package b2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type ErrorJSON struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func backblaze_error(errMsg ErrorJSON) error {
	return fmt.Errorf("%s: %s", errMsg.Code, errMsg.Message)
}

func Authorize(accountID, secretKey string) (*B2, error) {
	id_and_key := []byte(accountID + ":" + secretKey)
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var errMsg ErrorJSON
		if err := json.Unmarshal(data, &errMsg); err != nil {
			return nil, err
		}
		return nil, backblaze_error(errMsg)
	}

	var b2 B2
	if err := json.Unmarshal(data, &b2); err != nil {
		return nil, err
	}

	b2.client = client
	return &b2, nil
}
