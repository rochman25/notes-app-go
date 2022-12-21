package models

type Authentications struct {
	Token string `json:"token" gorm:"type:TEXT;not null"`
}
