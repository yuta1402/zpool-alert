package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type requestData struct {
	Text string `json:"text"`
}

func PostAlert(errorMessage string, apiURL string) (*http.Response, error) {
	text := "```" + errorMessage + "```"

	data := requestData{
		Text: text,
	}

	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer([]byte(json)))
	if err != nil {
		return nil, err
	}

	return res, nil
}
