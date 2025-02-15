# README.md

# Language Learning Portal API

## Overview

This project is a backend API for a language learning portal designed to facilitate vocabulary learning, track study sessions, and manage learning activities. The API is built using Go and utilizes the Gin framework for handling HTTP requests.

## Project Structure

- **cmd/server**: Contains the main entry point for the server application. Initializes the Gin router and sets up API endpoints.
- **internal/models**: Defines data structures and database operations related to vocabulary words, groups, study sessions, and review items.
- **internal/handlers**: Contains HTTP handlers organized by feature (dashboard, words, groups, etc.) to manage incoming requests and responses.
- **internal/service**: Implements the business logic of the application, including functions to interact with models and handle data processing.
- **db/migrations**: Contains SQL migration files to set up and modify the database schema as needed.
- **db/seeds**: Contains JSON seed files for populating the database with initial data, defining vocabulary words and their corresponding groups.
- **magefile.go**: Defines tasks for Mage, the task runner for Go, including tasks for initializing the database, running migrations, and seeding data.
- **go.mod**: Defines the module and its dependencies for the Go project.

## Setup Instructions

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd backend_go
   ```

2. **Install dependencies**:
   ```
   go mod tidy
   ```

3. **Initialize the database**:
   ```
   mage initialize
   ```

4. **Run migrations**:
   ```
   mage migrate
   ```

5. **Seed the database**:
   ```
   mage seed
   ```

6. **Start the server**:
   ```
   go run cmd/server/main.go
   ```

## API Usage

The API provides various endpoints for managing vocabulary words, study sessions, and groups. Refer to the API documentation for detailed information on available endpoints and their usage.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.