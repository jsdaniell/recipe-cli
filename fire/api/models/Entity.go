package models

// Entity model used to receive request and send JSON response, is used on authentication controller.
type Entity struct {
	ID            string   `json:"_id"`
	Name          string   `json:"name, omitempty"`
	Description   string   `json: "description"`
}
