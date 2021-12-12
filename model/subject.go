package model

type SubjectModel struct {
	ID      string        `gorm:"primaryKey"`
	Title   string        `gorm:"unique;not null"`
	Members []MemberModel `gorm:"foreignKey:ID"`
}
