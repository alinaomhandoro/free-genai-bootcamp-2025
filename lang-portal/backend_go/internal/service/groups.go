package service

import (
    "backend_go/internal/models"
    "gorm.io/gorm"
)

// GetGroups returns a list of groups
func GetGroups(db *gorm.DB) ([]models.Group, error) {
    var groups []models.Group
    if err := db.Find(&groups).Error; err != nil {
        return nil, err
    }
    return groups, nil
}
