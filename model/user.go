package model

import (
	"time"
)

type User struct {
	Id         int64     `gorm:"primary_key;auto_increment" json:"id"`
	Admin      bool      `gorm:"default:false" json:"admin"`
	Management bool      `gorm:"default:false" json:"management"`
	Name       string    `gorm:"type:varchar(10);unique;not null" json:"name"`
	Password   string    `gorm:"not null" json:"-"`
	Email      string    `gorm:"unique;not null" json:"email"`
	CreateAt   time.Time `gorm:"not null" json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

func (User) TableName() string {
	return "mock_user_t"
}
