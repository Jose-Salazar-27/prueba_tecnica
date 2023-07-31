package services

import (
	"time"

	"github.com/Jose-Salazar-27/prueba_tecnica/models"
	"github.com/Jose-Salazar-27/prueba_tecnica/repositories"
)

type UserService struct {
	Repository repositories.Repository
}

func NewUserService() (*UserService, error) {
	repo, err := repositories.NewRepository()

	if err != nil {
		return nil, err
	}

	return &UserService{
		Repository: repo,
	}, nil
}

var UserMock = []*models.User{
	&models.User{
		ID:         1,
		Name:       "jose",
		LastName:   "sala",
		Email:      "sample@email",
		Profession: "programmer",
		Created_at: time.Now().UTC(),
	},
	&models.User{
		ID:         2,
		Name:       "andres",
		LastName:   "sala",
		Email:      "sample@email",
		Profession: "programmer",
		Created_at: time.Now().UTC(),
	},
}

func (rec *UserService) GetUser(id int) (*models.User, error) {
	return UserMock[id], nil
}

func (rec *UserService) CreateUser(user *models.User) (*models.User, error) {
	record, err := rec.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	return record, nil
}
