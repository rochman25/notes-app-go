package models

import "time"

type Notes struct {
	ID        int       `json:"id,primary_key" gorm:"type:BIGINT"`
	Title     string    `json:"title" gorm:"type:text;not null"`
	Body      string    `json:"body" gorm:"type:text"`
	Tags      string    `json:"tags" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
