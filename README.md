# Task Management API

This is a RESTful Task Management API built with Go, using the Gin framework, SQLite for the database, and Docker for containerization. It allows users to create, retrieve, update, and delete tasks, with both public and protected endpoints. 

---

## Features
- **CRUD Operations**: Create, read, update, and delete tasks.
- **Public & Protected Endpoints**: Token-based authentication for secured endpoints.
- **Database**: SQLite for persistence, with migrations handled by GORM.
- **Pagination**: Task listing supports pagination.
- **Environment Configurations**: Configurations are stored in `.env` files for flexibility.
- **Dockerized**: Easily run the application in a containerized environment.

---

## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Docker](#docker)
- [Project Structure](#project-structure)

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mukeshmahajan/task-management-api.git
   cd task-management-api

2. **Install dependency**:
  go mod tidy

## Configuration

**Create a .env file in the root directory and add the following configurations**
  DB_NAME=task.db
  PORT=8080
  AUTH_TOKEN=your-secure-api-key

## Usage

**Run Migrations**:
Migrations are automatically executed when the application starts. If needed, you can run them manually by starting the application:
  go run main.go

**Start the Server**:
Start the application locally:
  go run main.go

## API Endpoints
  **Public Endpoints**:
  Method	 Endpoint	      Description
  GET	     /public/tasks	Retrieve all tasks
  **Protected Endpoints**:
  Method	 Endpoint	      Description
  GET	     /tasks/:id	    Retrieve a task by its ID
  POST	   /tasks/	      Create a new task
  PUT	     /tasks/:id   	Update an existing task
  DELETE	 /tasks/:id   	Delete a task by its ID

## Testing

**Run the unit and integration tests using**:
go test ./...

# Docker

**Build the Docker Image**
docker build -t task-management-api .

**Run the Container**
docker run -p 8080:8080 --env-file .env task-management-api

## Project Structure

task-management-api/
├── controllers/        # Endpoint handlers
│   ├── protected.go    # Handlers for protected task CRUD operations
|   ├── public.go       # Handlers for public task CRUD operations
├── database/           # Database connection and setup
│   ├── db.go           # SQLite connection logic
├── middleware/         # Authentication middleware
│   ├── auth.go         # API key validation logic
├── migrations/         # Database migrations
│   ├── migration.go    # Migration logic
├── models/             # Models definition
│   ├── task.go         # Task model
├── routes/             # API routes
│   ├── protected.go    # Protected Route definitions
|   ├── public.go       # Public Route definitions
├── tests/              # Authentication middleware
│   ├── auth.go         # Handlers for test operations
├── main.go             # Entry point of the application
├── .env                # Environment variables
├── .gitignore          # Git ignore file
├── .dockerignore       # docker ignore file
├── Dockerfile          # Docker setup
├── README.md           # Project documentation
└── go.mod              # Dependency management
