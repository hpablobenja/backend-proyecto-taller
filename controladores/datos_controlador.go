package controladores

import (
	"encoding/json"
	"log"
	"net/http"

	"crud-mysql/conexiones"
	"crud-mysql/modelos"

	"github.com/gorilla/mux"
)

func ObtenerUsuarios(writer http.ResponseWriter, request *http.Request) {
	personas := []modelos.Persona{}
	db := conexiones.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	conexiones.SendReponse(writer, http.StatusOK, json)
}

func ObtenerNoticias(writer http.ResponseWriter, request *http.Request) {
	noticias := []modelos.Noticia{}
	db := conexiones.GetConnection()
	defer db.Close()

	db.Find(&noticias)
	json, _ := json.Marshal(noticias)
	conexiones.SendReponse(writer, http.StatusOK, json)
}

func Encontrar(writer http.ResponseWriter, request *http.Request) {
	persona := modelos.Persona{}

	id := mux.Vars(request)["id"]

	db := conexiones.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		conexiones.SendReponse(writer, http.StatusOK, json)
	} else {
		conexiones.SendError(writer, http.StatusNotFound)
	}

}

func GuardarUsuario(writer http.ResponseWriter, request *http.Request) {
	persona := modelos.Persona{}

	db := conexiones.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		conexiones.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		conexiones.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	conexiones.SendReponse(writer, http.StatusCreated, json)
}
func GuardarNoticia(writer http.ResponseWriter, request *http.Request) {
	noticia := modelos.Noticia{}

	db := conexiones.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&noticia)

	if error != nil {
		log.Fatal(error)
		conexiones.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&noticia).Error

	if error != nil {
		log.Fatal(error)
		conexiones.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(noticia)

	conexiones.SendReponse(writer, http.StatusCreated, json)
}

func BorrarUsuario(writer http.ResponseWriter, request *http.Request) {
	persona := modelos.Persona{}

	db := conexiones.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		conexiones.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		conexiones.SendError(writer, http.StatusNotFound)
	}
}
func BorrarNoticia(writer http.ResponseWriter, request *http.Request) {
	noticia := modelos.Noticia{}

	db := conexiones.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&noticia, id)

	if noticia.ID > 0 {
		db.Delete(noticia)
		conexiones.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		conexiones.SendError(writer, http.StatusNotFound)
	}
}
