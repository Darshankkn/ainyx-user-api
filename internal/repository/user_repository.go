package repository

import (
	"context"
	

	db "ainyx-user-api/db/sqlc/generated"
	"ainyx-user-api/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB      *pgx.Conn
	Queries *db.Queries
}

func NewUserRepository(dbConn *pgx.Conn) *UserRepository {
	return &UserRepository{
		DB:      dbConn,
		Queries: db.New(dbConn),
	}
}
func mapSQLCUser(user db.User) models.User {
	return models.User{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob.Time,
	}
}

func (r *UserRepository) CreateUser(user models.CreateUserRequest) (models.User, error) {

	var createdUser models.User

	query := `
		INSERT INTO users(name, dob)
		VALUES($1, $2)
		RETURNING id, name, dob
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.DOB,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.DOB,
	)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (r *UserRepository) GetUserByID(id int) (models.User, error) {

	user, err := r.Queries.GetUserByID(
		context.Background(),
		int32(id),
	)

	if err != nil {
		return models.User{}, err
	}

	return mapSQLCUser(user), nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {

	query := `
		SELECT id, name, dob
		FROM users
		ORDER BY id
	`

	rows, err := r.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.DOB,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(id int, user models.UpdateUserRequest) (models.User, error) {

	var updatedUser models.User

	query := `
		UPDATE users
		SET name = $1,
		    dob = $2
		WHERE id = $3
		RETURNING id, name, dob
	`

	err := r.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.DOB,
		id,
	).Scan(
		&updatedUser.ID,
		&updatedUser.Name,
		&updatedUser.DOB,
	)

	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *UserRepository) DeleteUser(id int) error {

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}