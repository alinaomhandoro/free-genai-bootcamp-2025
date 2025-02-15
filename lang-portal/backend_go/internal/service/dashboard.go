package service

import (
    "backend_go/internal/models"
    "gorm.io/gorm"
)

// GetLastStudySession returns the most recent study session
func GetLastStudySession(db *gorm.DB) (*models.StudySession, error) {
    var session models.StudySession
    if err := db.Order("created_at desc").First(&session).Error; err != nil {
        return nil, err
    }
    return &session, nil
}
