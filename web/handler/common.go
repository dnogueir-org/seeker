package handler

import "encoding/json"

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}

func responseBodyToJSON(body interface{}) ([]byte, error) {
	responseJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return responseJSON, nil
}
