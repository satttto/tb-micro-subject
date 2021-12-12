package model

type MemberModel struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}
