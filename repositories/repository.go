package repositories

import "github.com/Jose-Salazar-27/prueba_tecnica/models"

type Repository interface {
	List() ([]models.User, error)
	FindById(id int) (*models.User, error)
	DeleteById(id int) error
	create() error
}
