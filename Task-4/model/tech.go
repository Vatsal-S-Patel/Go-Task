package model

type Technology struct {
	Id          int          `json:"id"`
	TechDetails []TechDetail `json:"techDets"`
}
