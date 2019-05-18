package models

import "encoding/json"

type RequestResult struct {
	Status  string          `json:"status"`
	Error   string          `json:"error"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
