package service

import (
    "backend_go/internal/models"
    "gorm.io/gorm"
)

// GetWords returns a list of words
func GetWords(db *gorm.DB) ([]models.Word, error) {
    var words []models.Word
    if err := db.Find(&words).Error; err != nil {
        return nil, err
    }
    return words, nil
}
