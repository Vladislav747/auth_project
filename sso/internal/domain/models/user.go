package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	PassHash []byte `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}
