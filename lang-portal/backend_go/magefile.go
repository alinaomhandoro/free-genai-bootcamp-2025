// magefile.go
// This file defines tasks for Mage, the task runner for Go.

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// InitializeDatabase initializes the SQLite database called words.db.
func InitializeDatabase() error {
	fmt.Println("Initializing database...")
	// Add logic to initialize the database
	return nil
}

// MigrateDatabase runs a series of migration SQL files on the database.
func MigrateDatabase() error {
	fmt.Println("Running database migrations...")
	migrationsDir := "db/migrations"
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := sh.Run("sqlite3", "words.db", ".read "+migrationsDir+"/"+file.Name()); err != nil {
			return err
		}
	}
	return nil
}

// SeedData imports JSON seed files and transforms them into target data for the database.
func SeedData() error {
	fmt.Println("Seeding data...")
	// Add logic to seed data from JSON files
	return nil
}

// Default task to run when no task is specified.
func Default() {
	mg.SerialDeps(InitializeDatabase, MigrateDatabase, SeedData)
}