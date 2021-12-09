package db

type MemberModel struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type SubjectModel struct {
	ID      string        `gorm:"primaryKey"`
	Title   string        `gorm:"unique;not null"`
	Members []MemberModel `gorm:"foreignKey:ID"`
}
