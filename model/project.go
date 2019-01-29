package model

import "time"

type Project struct {
	Id                 int64     `gorm:"primary_key;auto_increment" json:"id"`
	Key                string    `gorm:"unique;not null" json:"key"`
	Name               string    `gorm:"type:varchar(20);unique;not null" json:"name"`
	UserId             int64     `gorm:"not null" json:"userId"`
	Host               string    `gorm:"type:varchar(255);not null" json:"host"`
	InsecureSkipVerify bool      `gorm:"default:false" json:"insecureSkipVerify"`
	Enable             bool      `gorm:"default:false" json:"enable"`
	CreatedAt          time.Time `gorm:"not null" json:"createdAt"`
	UpdateAt           time.Time `json:"updateAt"`
}

func (Project) TableName() string {
	return "mock_project_t"
}
