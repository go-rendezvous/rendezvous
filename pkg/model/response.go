package model

type Response struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value,omitempty"`
}
