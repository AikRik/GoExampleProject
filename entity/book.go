package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(32)" json:"firstname"`
	Lastname  string `gorm:"type:varchar(32)" json:"lastname"`
	Age       int8   `binding:"gte=1, lte=130" json:"age"`
}

type Book struct {
	ID          uint64    `gorm:"primary_key;auto_increment" "json:"id"`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:type:"varchar(100);UNIQUE"`
	Description string    `json:"title" binding:"min=2,max=200" gorm:type:"varchar(200)"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	Quantity    int       `json:"quantity"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
