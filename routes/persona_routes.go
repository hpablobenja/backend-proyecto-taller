package routes

import (
	"crud-mysql/controllers"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/noticias", controllers.GetNoticias).Methods("GET")
	subRoute.HandleFunc("/nsave", controllers.Nsave).Methods("POST")
	subRoute.HandleFunc("/ndelete/{id}", controllers.Ndelete).Methods("POST")
}
