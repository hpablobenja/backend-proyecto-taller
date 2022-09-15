package main

import (
	"log"
	"net/http"
	"os"
	"crud-mysql/commons"
	"crud-mysql/routes"

	"github.com/gorilla/mux"
)
//Addr:    ":9000",
func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	commons.EnableCORS(router)

	server := http.Server{
		Addr:  os.Getenv("PORT"),
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServe())
}
