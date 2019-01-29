package model

import "time"

type Data struct {
	Id           int64     `gorm:"primary_key;auto_increment" json:"id"`
	RuleId       int64     `gorm:"not null" json:"ruleId"`
	UserId       int64     `gorm:"not null" json:"userId"`
	ResponseCode int       `gorm:"not null;default:200" json:"responseCode"`
	ContentType  string    `gorm:"not null;default:json" json:"contentType"` // text html json xml
	Content      string    `gorm:"type:text;not null" json:"content"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdateAt     time.Time `json:"updateAt"`
}

func (Data) TableName() string {
	return "mock_data_t"
}
