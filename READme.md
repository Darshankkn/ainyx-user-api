# Ainyx User API

A RESTful User Management API built using GoFiber, PostgreSQL, SQLC, Uber Zap Logger, and go-playground/validator.

## Features

* Create User
* Get User By ID
* Get All Users
* Update User
* Delete User
* Dynamic Age Calculation from DOB
* Input Validation
* Structured Logging with Uber Zap
* SQLC Generated Database Layer

## Tech Stack

* Go
* GoFiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator
* pgx

## Project Structure

cmd/server/main.go

config/

db/

* migrations/
* sqlc/

  * schema.sql
  * query.sql
  * generated/

internal/

* handler/
* repository/
* service/
* routes/
* middleware/
* models/
* logger/

## Database Schema

users

| Column | Type               |
| ------ | ------------------ |
| id     | SERIAL PRIMARY KEY |
| name   | TEXT               |
| dob    | DATE               |

## Setup

### Clone Repository

git clone <repository-url>

cd ainyx-user-api

### Create Database

CREATE DATABASE ainyx_users;

### Create Table

CREATE TABLE users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
dob DATE NOT NULL
);

### Run Application

go mod tidy

go run cmd/server/main.go

Server starts on:

http://localhost:3000

## API Endpoints

POST /users

GET /users

GET /users/:id

PUT /users/:id

DELETE /users/:id

## Dynamic Age Calculation

Age is not stored in the database. Only the date of birth is stored. Age is calculated dynamically using Go's time package whenever user data is fetched.

## SQLC

SQL queries are defined in query.sql and SQLC generates type-safe Go code for database access.

## Logging

Uber Zap is used for structured logging of user operations.

## Validation

go-playground/validator is used to validate request payloads before processing.
