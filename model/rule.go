package model

import "time"

type Rule struct {
	Id        int64     `gorm:"primary_key;auto_increment" json:"id"`
	ProjectId int64     `gorm:"not null" json:"projectId"`
	Path      string    `gorm:"not null" json:"path"`
	Method    string    `gorm:"not null" json:"method"`
	Enable    bool      `gorm:"default:false" json:"enable"`
	CreatedAt time.Time `gorm:"not null" json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

func (Rule) TableName() string {
	return "mock_rule_t"
}
