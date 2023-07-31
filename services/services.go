package services

import (
	"github.com/Jose-Salazar-27/prueba_tecnica/common"
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

func (rec *UserService) GetUser(id int) (*models.User, error) {

	result, err := rec.Repository.FindById(id)

	if id == 0 {
		err := common.Exception{Message: "id invalid or not provided"}
		return nil, &err
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *UserService) CreateUser(user *models.User) (*models.User, error) {
	record, err := rec.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func (rec *UserService) GetAllUsers(users *[]models.User) (any, error) {
	// var users []*models.User

	result, err := rec.Repository.List(users)

	if err != nil {
		return nil, err
	}

	return result, nil

}
