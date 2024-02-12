package model

type FinalOutput struct {
	Userid      int
	Name        string
	UserAddress Address `json:"Address"`
	TechDetails []NewTechDetail
	Email       string
	Phone       string
}
