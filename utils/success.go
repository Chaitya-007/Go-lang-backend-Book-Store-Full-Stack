package utils

import (
	"bookstore/schema"
	"encoding/json"
	"net/http"
)

// # Send Success Response
func SendSuccess(w http.ResponseWriter, msg interface{}) {
	// # Set Headers for Success Response
	w.WriteHeader(http.StatusOK) // # 200

	// # Create New Encoder for Response Body
	encoder := json.NewEncoder(w)

	// # Create Success Response Object
	res := &schema.Success{
		Success: true,
		Message: msg,
	}

	// # Encode Response to Writer
	encoder.Encode(res)
}
