package model

// Book struct
type Book struct {
	// gorm.Model
	Id        int    `json:"id" gorm:"column:id;primaryKey"`
	Title     string `json:"title" gorm:"column:title"`
	Author    string `json:"author" gorm:"column:author"`
	ISBN      string `json:"isbn" gorm:"column:isbn"`
	Publisher string `json:"publisher" gorm:"column:publisher"`
	Year      int16  `json:"year" gorm:"column:year"`
	Genre     string `json:"genre" gorm:"column:genre"`
}
