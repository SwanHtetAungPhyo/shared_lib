package models

type Email struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}
