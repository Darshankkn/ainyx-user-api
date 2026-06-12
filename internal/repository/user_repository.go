package repository

import (
	"context"
	"time"

	db "ainyx-user-api/db/sqlc/generated"
	"ainyx-user-api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

	parsedTime, err := time.Parse("2006-01-02", user.DOB)

if err != nil {
	return models.User{}, err
}

dob := pgtype.Date{
	Time:  parsedTime,
	Valid: true,
}

	if err != nil {
		return models.User{}, err
	}

	createdUser, err := r.Queries.CreateUser(
		context.Background(),
		db.CreateUserParams{
			Name: user.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		return models.User{}, err
	}

	return mapSQLCUser(createdUser), nil
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

	users, err := r.Queries.GetAllUsers(context.Background())

	if err != nil {
		return nil, err
	}

	var result []models.User

	for _, user := range users {
		result = append(result, mapSQLCUser(user))
	}

	return result, nil
}

func (r *UserRepository) UpdateUser(id int, user models.UpdateUserRequest) (models.User, error) {

	parsedTime, err := time.Parse("2006-01-02", user.DOB)

if err != nil {
	return models.User{}, err
}

dob := pgtype.Date{
	Time:  parsedTime,
	Valid: true,
}

	if err != nil {
		return models.User{}, err
	}

	updatedUser, err := r.Queries.UpdateUser(
		context.Background(),
		db.UpdateUserParams{
			Name: user.Name,
			Dob:  dob,
			ID:   int32(id),
		},
	)

	if err != nil {
		return models.User{}, err
	}

	return mapSQLCUser(updatedUser), nil
}
func (r *UserRepository) DeleteUser(id int) error {

	return r.Queries.DeleteUser(
		context.Background(),
		int32(id),
	)
}