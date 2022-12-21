package models

import "time"

type Users struct {
	ID        int       `json:"id,primary_key" gorm:"type:BIGINT"`
	Username  string    `json:"username" gorm:"type:VARCHAR(50);not null;unique"`
	Password  string    `json:"password" gorm:"type:TEXT;not null"`
	FullName  string    `json:"fullname" gorm:"type:TEXT;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
