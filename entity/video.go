package entity

import (
	"time"
)

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id" `
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=120"`
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(128);"`
}

func (Person) TableName() string {
	return "person"
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=20" validate:"is-cool" gorm:"type:varchar(20)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url,max=256" gorm:"type:varchar(256);UNIQUE"`
	Author      *Person   `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-" gorm:"index; no null"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
