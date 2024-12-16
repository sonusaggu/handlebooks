package models

type Book struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"column:title"`
	Author string `json:"author" gorm:"column:author"`
}
