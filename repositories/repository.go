package repositories

import (
	"fmt"

	"github.com/Jose-Salazar-27/prueba_tecnica/common"
	"github.com/Jose-Salazar-27/prueba_tecnica/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	List(entity *[]models.User) (any, error) //done
	FindById(id int) (*models.User, error)   //done
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

func (rec *PostgresRepository) List(entity *[]models.User) (any, error) {

	result := rec.connector.Find(&entity)
	fmt.Println(result)

	if result.Error != nil {
		return nil, result.Error
	}

	return result, nil
}

func (rec *PostgresRepository) FindById(id int) (*models.User, error) {

	var user models.User
	rec.connector.First(&user, id)

	if user.ID == 0 {
		error := common.Exception{Message: "user not found"}
		return nil, &error
	}

	return &user, nil
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
