package service

import (
	"time"

	"ainyx-user-api/internal/models"
	"ainyx-user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user models.CreateUserRequest) (models.User, error) {
	return s.repo.CreateUser(user)
}

func calculateAge(dob time.Time) int {

	now := time.Now()

	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func (s *UserService) GetUserByID(id int) (models.UserResponse, error) {

	user, err := s.repo.GetUserByID(id)

	if err != nil {
		return models.UserResponse{}, err
	}

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB.Format("2006-01-02"),
		Age:  calculateAge(user.DOB),
	}

	return response, nil
}

func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {

	users, err := s.repo.GetAllUsers()

	if err != nil {
		return nil, err
	}

	var response []models.UserResponse

	for _, user := range users {

		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.DOB.Format("2006-01-02"),
			Age:  calculateAge(user.DOB),
		})
	}

	return response, nil
}

func (s *UserService) UpdateUser(id int, user models.UpdateUserRequest) (models.User, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}