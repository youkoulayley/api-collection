package models

// JSONError is the type for json errors
type JSONError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
