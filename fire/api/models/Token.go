package models

// User model used to receive request and send JSON response, is used on authentication controller.
type Token struct {
	CreatedAt string `json:"created_at,omitempty"`
	Token     string `json:"token,omitempty"`
}
