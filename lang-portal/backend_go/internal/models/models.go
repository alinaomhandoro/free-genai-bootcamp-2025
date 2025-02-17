package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Word represents a vocabulary word
type Word struct {
	ID       uint           `gorm:"primaryKey"`
	Japanese string         `json:"japanese"`
	Romaji   string         `json:"romaji"`
	English  string         `json:"english"`
	Parts    datatypes.JSON `json:"parts"`
}

// Group represents a thematic group of words
type Group struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}

// StudySession represents a study session
type StudySession struct {
	ID              uint      `gorm:"primaryKey"`
	GroupID         uint      `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID uint      `json:"study_activity_id"`
}

// StudyActivity represents a specific study activity
type StudyActivity struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

// WordReviewItem represents a record of word practice
type WordReviewItem struct {
	WordID    uint      `json:"word_id"`
	CreatedAt time.Time `json:"created_at"`
	Correct   bool      `json:"correct"`
}

// InitializeDatabase initializes the SQLite database
func InitializeDatabase(db *gorm.DB) {
	db.AutoMigrate(&Word{}, &Group{}, &StudySession{}, &StudyActivity{}, &WordReviewItem{})
}

// CreateWord creates a new word in the database
func CreateWord(db *gorm.DB, word *Word) error {
	return db.Create(word).Error
}

// GetWord retrieves a word by ID from the database
func GetWord(db *gorm.DB, id uint) (*Word, error) {
	var word Word
	if err := db.First(&word, id).Error; err != nil {
		return nil, err
	}
	return &word, nil
}

// UpdateWord updates an existing word in the database
func UpdateWord(db *gorm.DB, word *Word) error {
	return db.Save(word).Error
}

// DeleteWord deletes a word by ID from the database
func DeleteWord(db *gorm.DB, id uint) error {
	return db.Delete(&Word{}, id).Error
}

// Similar methods can be created for Group, StudySession, StudyActivity, and WordReviewItem
