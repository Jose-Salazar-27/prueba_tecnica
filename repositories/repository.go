package repositories

import (
	"time"

	"github.com/Jose-Salazar-27/prueba_tecnica/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	List() ([]*models.User, error)         //done
	FindById(id int) (*models.User, error) //done
	DeleteById(id int) error
	Create(user *models.User) (*models.User, error)
}

type PostgresRepository struct {
	connector *gorm.DB
}

func NewRepository() (*PostgresRepository, error) {
	dsn := "host=localhost user=root password=psw1234 dbname=jose_db port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	migrationErr := connection.AutoMigrate(models.User{})

	if migrationErr != nil {
		return nil, err
	}

	return &PostgresRepository{
		connector: connection,
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

func (rec *PostgresRepository) List() ([]*models.User, error) {
	return UserMock, nil
}

func (rec *PostgresRepository) FindById(id int) (*models.User, error) {
	return UserMock[0], nil
}

func (rec *PostgresRepository) DeleteById(id int) error {
	return nil
}

func (rec *PostgresRepository) Create(user *models.User) (*models.User, error) {
	result := rec.connector.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
