# Ainyx User API

A RESTful User Management API built using **GoFiber**, **PostgreSQL**, **SQLC**, **Uber Zap Logger**, and **go-playground/validator**.

The API allows users to be created, retrieved, updated, and deleted while dynamically calculating age from the stored date of birth.

---

## Features

### Core Features

* Create User
* Get User By ID
* Get All Users
* Update User
* Delete User
* Dynamic Age Calculation from Date of Birth (DOB)
* Input Validation using go-playground/validator
* Structured Logging using Uber Zap
* SQLC Generated Database Access Layer
* PostgreSQL Integration
* Layered Architecture (Handler → Service → Repository)

### Bonus Features Implemented

* Request ID Middleware
* Request Duration Logging Middleware

---

## Tech Stack

* Go
* GoFiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator
* pgx

---

## Project Structure

```text
cmd/
└── server/
    └── main.go

config/

db/
├── migrations/
└── sqlc/
    ├── schema.sql
    ├── query.sql
    └── generated/

internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

---

## Database Schema

### users

| Column | Type   | Constraints |
| ------ | ------ | ----------- |
| id     | SERIAL | PRIMARY KEY |
| name   | TEXT   | NOT NULL    |
| dob    | DATE   | NOT NULL    |

---

## Setup Instructions

### 1. Clone Repository

```bash
git clone <repository-url>
cd ainyx-user-api
```

### 2. Create Database

```sql
CREATE DATABASE ainyx_users;
```

### 3. Create Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run Application

```bash
go run cmd/server/main.go
```

Server starts on:

```text
http://localhost:3000
```

---

## API Endpoints

| Method | Endpoint   | Description    |
| ------ | ---------- | -------------- |
| POST   | /users     | Create User    |
| GET    | /users     | Get All Users  |
| GET    | /users/:id | Get User By ID |
| PUT    | /users/:id | Update User    |
| DELETE | /users/:id | Delete User    |

---

## Sample Request

### Create User

```json
{
  "name": "Darshan",
  "dob": "2001-07-15"
}
```

### Response

```json
{
  "id": 1,
  "name": "Darshan",
  "dob": "2001-07-15"
}
```

---

## Dynamic Age Calculation

Age is not stored in the database.

Only the user's date of birth is stored. Age is calculated dynamically using Go's `time` package whenever user information is retrieved.

---

## SQLC

SQL queries are defined in:

```text
db/sqlc/query.sql
```

SQLC generates type-safe Go code which is used by the repository layer for database operations.

---

## Logging

Uber Zap is used for structured logging of:

* User Creation
* User Updates
* User Deletion
* Request Processing Duration

Example:

```json
{
  "msg": "Request Completed",
  "method": "GET",
  "path": "/users",
  "duration": "1.2ms"
}
```

---

## Request ID Middleware

Every request is assigned a unique Request ID.

Example Response Header:

```text
X-Request-ID: 5741c401-dd23-4185-a3c7-587749059ac8
```

---

## Validation

go-playground/validator is used to validate incoming request payloads before processing API requests.

---

## Author

Darshan K N
