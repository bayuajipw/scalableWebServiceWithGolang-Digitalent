package model

import "time"

// Account model info
// @Description User account information
// @Description with user id, fullname, and datebirth
type Account struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"full_name"`
	Datebirth time.Time `json:"date_birth"`
}
