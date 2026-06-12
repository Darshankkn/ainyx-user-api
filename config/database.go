package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {

	connStr := "postgres://postgres:Darshan@511@localhost:5432/ainyx_users"

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully")

	return conn, nil
}