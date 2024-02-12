package model

type Contact struct {
	Id          int           `json:"id"`
	UserContact ContactDetail `json:"contactDets"`
}
