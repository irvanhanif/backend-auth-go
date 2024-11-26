package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 			uuid.UUID	
	Fullname 	string 			`gorm:"type:varchar(100);not null" json:"fullname" binding:"required,min=5"`
	Username 	string 			`gorm:"type:varchar(50);unique;not null" json:"username" binding:"required,min=4"`
	Email 		string 			`gorm:"type:varchar(50);unique;not null" json:"email" binding:"required,min=5,email"`
	Password 	string 			`gorm:"type:varchar(150);not null" json:"password" binding:"required,min=6"`
	Born 		*string 		`gorm:"type:date" json:"born" time_format:"2024-11-20"`
	Gender 		*string 		`gorm:"type:boolean" json:"gender"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type UserLogin struct {
	EmailOrUsername string `json:"emailOrUsername" binding:"required,min=5"`
	Password		string `json:"password" binding:"required,min=6"`
}