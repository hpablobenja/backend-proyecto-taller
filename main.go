package main

import (
	"log"
	"net/http"
	//"os"
	"crud-mysql/conexiones"
	"crud-mysql/rutas"

	"github.com/gorilla/mux"
)
//Addr:    ":9000",
//Addr:  ":"+port
func main() {
	conexiones.Migrate()

	router := mux.NewRouter()
	rutas.SetPersonaRoutes(router)
	conexiones.EnableCORS(router)
	//port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}
