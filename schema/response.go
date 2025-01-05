package schema

// # Success Response Schema
type Success struct {
	Success bool        `json:"success"` // # true
	Message interface{} `json:"message"`
}

// # Error Response Schema
type Error struct {
	Success bool        `json:"success"` // # false
	Error   interface{} `json:"error"`
}
