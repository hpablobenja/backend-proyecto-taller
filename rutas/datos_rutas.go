package rutas

import (
	"crud-mysql/controladores"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/datos/api").Subrouter()
	subRoute.HandleFunc("/personas", controladores.ObtenerUsuarios).Methods("GET")
	subRoute.HandleFunc("/pguardar", controladores.GuardarUsuario).Methods("POST")
	subRoute.HandleFunc("/pborrar/{id}", controladores.BorrarUsuario).Methods("POST")
	subRoute.HandleFunc("/encontrar/{id}", controladores.Encontrar).Methods("GET")
	subRoute.HandleFunc("/noticias", controladores.ObtenerNoticias).Methods("GET")
	subRoute.HandleFunc("/nguardar", controladores.GuardarNoticia).Methods("POST")
	subRoute.HandleFunc("/nborrar/{id}", controladores.BorrarNoticia).Methods("POST")
}
