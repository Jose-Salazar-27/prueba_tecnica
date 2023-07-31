package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         int       `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(50)"`
	LastName   string    `gorm:"type:varchar(50)"`
	Email      string    `gorm:"unique_index"`
	Profession string    `gorm:"type:varchar(50)"`
	Created_at time.Time `gorm:"autoCreateTime:true"`
}
