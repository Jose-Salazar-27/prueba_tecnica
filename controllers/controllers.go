package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jose-Salazar-27/prueba_tecnica/common"
	"github.com/Jose-Salazar-27/prueba_tecnica/models"
	"github.com/Jose-Salazar-27/prueba_tecnica/services"
	"github.com/gorilla/mux"
)

type Controller func(http.ResponseWriter, *http.Request) error

type UserController struct {
	service services.UserService
}

func NewUserController() (*UserController, error) {
	service, err := services.NewUserService()

	if err != nil {
		return nil, err
	}

	return &UserController{
		service: *service,
	}, nil
}

func (rec *UserController) HandleUserRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return rec.ListUsers(w, r)
	}

	if r.Method == http.MethodPost {
		return rec.Create(w, r)
	}

	if r.Method == http.MethodDelete {
		return rec.DeleteById(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (rec *UserController) UserByIdRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return rec.GetUserById(w, r)
	}

	if r.Method == http.MethodDelete {
		return rec.DeleteById(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// TODO colocar la funcionalidad que correponde
func (rec *UserController) ListUsers(w http.ResponseWriter, r *http.Request) error {
	var users []models.User
	_, err := rec.service.GetAllUsers(&users)

	if err != nil {
		common.JsonWriter(w, http.StatusBadRequest, err)
	}

	common.JsonWriter(w, http.StatusOK, users)
	return nil
}

func (rec *UserController) Create(w http.ResponseWriter, r *http.Request) error {
	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)

	result, err := rec.service.CreateUser(user)

	if err != nil {
		return err
	}

	common.JsonWriter(w, http.StatusCreated, result)
	return nil
}

func (rec *UserController) GetUserById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println(id)

	result, err := rec.service.GetUser(id)

	if err != nil {
		return err
	}

	common.JsonWriter(w, http.StatusCreated, result)
	return nil
}

func (rec *UserController) DeleteById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println(id)

	err := rec.service.DeleteUserById(id)

	if err != nil {
		return common.JsonWriter(w, http.StatusBadRequest, err)
	}

	return common.JsonWriter(w, http.StatusNoContent, nil)
}
