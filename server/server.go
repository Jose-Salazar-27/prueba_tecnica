package server

import (
	"net/http"

	"github.com/Jose-Salazar-27/prueba_tecnica/common"
	"github.com/Jose-Salazar-27/prueba_tecnica/controllers"
	"github.com/gorilla/mux"
)

type Server struct {
	listenAdrr string
}

func NewServer(port string) *Server {
	return &Server{
		listenAdrr: port,
	}
}

func (rec *Server) Init() {
	router := mux.NewRouter()
	userContoller, err := controllers.NewUserController()
	if err != nil {
		panic("Cannot initialze controller properly")
	}

	router.HandleFunc("/users", rec.MakeHandler(userContoller.HandleRequest))
	// router.HandleFunc("/users", rec.MakeHandler(userContoller.Create)).Methods(http.MethodPost)

	http.ListenAndServe(rec.listenAdrr, router)
}

func (rec *Server) MakeHandler(fn controllers.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			common.JsonWriter(w, http.StatusInternalServerError, err)
		}
	}
}
