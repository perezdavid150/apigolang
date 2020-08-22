package main

/*
Importación de librerias
*/
import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)

/*
Base de datos local en formato JSON
*/

type contacto struct{
	ID int `json:ID`
	Nombre string `json:Nombre`
	Telefono string `json:Telefono`
}

/*
Arreglo de base de datos
*/
type allContactos []contacto

var contactos = allContactos{
	{
		ID: 1,
		Nombre: "Juan",
		Telefono: "123456",
	},
}

/*
Método GET para obtener todos mis contactos
*/

func getContactos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactos)
}

/*
Método POST para crear un contacto nuevo
*/

func createContacto(w http.ResponseWriter, r *http.Request){
	var newContacto contacto
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Este contacto no es valido")
	}

	json.Unmarshal(reqBody, &newContacto)

	newContacto.ID = len(contactos) + 1

	contactos = append(contactos, newContacto)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newContacto)

}

/*
Método GET con para obtener 1 solo registro
*/

func getContacto(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	contactoID, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Fprintf(w, "ID no válido")
		return
	}

	for _, contacto := range contactos {
		if contacto.ID == contactoID{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(contacto)
		}
	}

}

/*
Método PUT para actualizar contacto
*/

func updateContacto(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	contactoID, err := strconv.Atoi(vars["id"])

	var updatedContacto contacto

	if err != nil{
		fmt.Fprintf(w, "ID no válido")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Datos no validos")
	}

	json.Unmarshal(reqBody, &updatedContacto)

	for i, contacto := range contactos {
		if contacto.ID == contactoID{
			contactos = append(contactos[:i], contactos[i + 1:]...)
			updatedContacto.ID = contactoID
			contactos = append(contactos, updatedContacto)

			fmt.Fprintf(w, "El contacto con ID %v fue actualizado", contactoID)

		}
	}

}

/*
Método DELETE para borrar contacto
*/

func deleteContacto(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	contactoID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID no válido")
		return
	}

	for i, contacto := range contactos {
		if contacto.ID == contactoID{
			contactos = append(contactos[:i], contactos[i + 1:]...)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "El contacto con ID %v fue removido", contactoID)
		}
	}
}

/*
Ruta base de la API
*/

func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ruta base")
}

/*
Definición de rutas
*/

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/contactos", getContactos).Methods("GET")
	router.HandleFunc("/contactos", createContacto).Methods("POST")
	router.HandleFunc("/contactos/{id}", getContacto).Methods("GET")
	router.HandleFunc("/contactos/{id}", deleteContacto).Methods("DELETE")
	router.HandleFunc("/contactos/{id}", updateContacto).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}