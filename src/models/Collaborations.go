package models

type Collaborations struct {
	ID     int `json:"id,primary_key" gorm:"type:BIGINT"`
	NoteID int
	Notes  Notes `gorm:"foreignKey:NoteID;constraint:OnDelete:CASCADE"`
	UserID int
	Users  Users `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
