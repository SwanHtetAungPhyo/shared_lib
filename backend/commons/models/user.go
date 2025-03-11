package models

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status,omitempty"`
}
