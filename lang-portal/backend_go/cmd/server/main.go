package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "backend_go/internal/models"
    "backend_go/internal/handlers"
    "backend_go/db/seeds"
)

func main() {
    // Initialize the database
    db, err := gorm.Open(sqlite.Open("words.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    // Migrate the schema
    models.InitializeDatabase(db)

    // Seed the database
    seeds.SeedDatabase(db)

    // Initialize Gin router
    r := gin.Default()

    // Middleware to provide the database connection to handlers
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    // Enable more detailed logging
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Define routes
    r.GET("/api/dashboard/last_study_session", handlers.GetLastStudySession)
    r.GET("/api/words", handlers.GetWords)
    r.GET("/api/groups", handlers.GetGroups)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
