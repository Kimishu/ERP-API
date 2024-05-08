package models

type Enterprise struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Sub   Subscription `json:"subscription"`
}
