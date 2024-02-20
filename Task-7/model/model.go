package model

// Profile Struct
type Profile struct {
	Id       uint   `json:"id" gorm:"column:id;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

// Task Struct
type Task struct {
	Id        uint   `json:"id" gorm:"column:id;primaryKey"`
	Title     string `json:"title" gorm:"column:title"`
	Body      string `json:"body" gorm:"column:body"`
	ProfileId uint   `json:"profile_id" gorm:"column:profile_id;constraint:OnDelete:CASCADE"`
}
