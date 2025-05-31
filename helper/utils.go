package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ParsePayload(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&target)
}

func ToJsonBody(v interface{}) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(v)

	if err != nil {
		return nil, err
	}

	return &buff, nil
}
