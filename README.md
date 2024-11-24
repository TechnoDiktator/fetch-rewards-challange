# fetch-rewards-challange
Assignment For The Fetch Rewards Challange (RECEIPT PROCESSOR)

# Receipt Processor API - Assignment Documentation

## Overview

This repository contains the solution for the **Receipt Processor** assignment. The task is to build a RESTful web service in Go that processes receipts and calculates points based on certain rules. The solution is implemented using the Gin framework for the HTTP server and an in-memory data store.

---

## Steps Taken

### 1. **Project Setup**
   - Initialized a new Go project using Go Modules.
   - Created a repository structure to support scalable development, following best practices for organizing Go code.
     - Set up folders for `cmd/`, `internal/`, `pkg/`, `config/`, and `tests/`.
     - Added a `Makefile` for automation of common tasks like running the server and running tests.

### 2. **Dependencies**
   - Installed Gin for the HTTP framework:
     ```bash
     go get -u github.com/gin-gonic/gin
     ```
   - Installed UUID for generating unique receipt IDs:
     ```bash
     go get github.com/google/uuid
     ```
   - (Optional) Installed Swagger for API documentation:
     ```bash
     go get -u github.com/swaggo/gin-swagger
     go get -u github.com/swaggo/files
     ```
   - (Optional) Installed Testify for unit testing:
     ```bash
     go get github.com/stretchr/testify
     ```

### 3. **Folder Structure**
   - Organized the code into logical modules:
     - `cmd/api/main.go` — Entry point for the application.
     - `internal/handlers/` — Contains the logic for handling HTTP requests.
     - `internal/services/` — Business logic for calculating points based on receipt data.
     - `internal/repository/` — In-memory store to hold receipts.
     - `internal/models/` — Structs and validation for receipt data.
     - `pkg/logger/` — Custom logger for the application.
     - `tests/` — Unit tests for the handlers and services.
   - Used `Makefile` for easy automation of common tasks.

### 4. **API Implementation**
   - Implemented two main endpoints:
     1. **`POST /receipts/process`**: Accepts a receipt in JSON format and generates a unique ID for it.
     2. **`GET /receipts/{id}/points`**: Retrieves the points for a receipt by ID.

### 5. **Points Calculation Logic**
   - Implemented the logic for calculating points based on the rules provided in the assignment.
     - Points for retailer name length.
     - Points for total amount based on the conditions (round dollar, multiple of 0.25, etc.).
     - Points for items based on description length and price.
     - Points for the day and time of purchase.

### 6. **Testing**
   - Wrote unit tests for both handlers and services.
   - Focused on testing the API response as well as the correctness of point calculations.

### 7. **Swagger Documentation (Optional)**
   - Added Swagger for API documentation to easily visualize and interact with the API.
   - The Swagger UI can be accessed by visiting `/swagger/index.html` (if using Swagger setup).

### 8. **Future Improvements**
   - Persisting data in a database (optional as per assignment).
   - Adding more unit tests for edge cases.
   - Enhancing logging and error handling.

---

## Running the Project

### Prerequisites

Make sure you have Go and Docker installed.

### 1. **Clone the Repository**
```bash
git clone https://github.com/TechnoDiktator/fetch-rewards-challange.git
cd fetch-rewards-challange

