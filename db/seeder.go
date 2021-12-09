package db

import (
	"log"

	"gorm.io/gorm"
)

func ExecuteSeeding(db *gorm.DB) {
	members1 := []MemberModel{
		{ID: "id1", Name: "Name1"},
		{ID: "id2", Name: "Name2"},
	}
	if err := db.Create(&members1).Error; err != nil {
		log.Fatalf("Failed to create members: %v", err)
	}
	members2 := []MemberModel{
		{ID: "id3", Name: "Name3"},
		{ID: "id4", Name: "Name4"},
	}
	if err := db.Create(&members2).Error; err != nil {
		log.Fatalf("Failed to create members: %v", err)
	}

	subjects := []SubjectModel{
		{ID: "sid1", Title: "title1", Members: members1},
		{ID: "sid2", Title: "title2", Members: members2},
	}
	if err := db.Create(&subjects).Error; err != nil {
		log.Fatalf("Failed to create subjects: %v", err)
	}
}
