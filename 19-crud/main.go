package main

import (
	"fmt"
	"log"
	"modulo/19-crud/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", service.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", service.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", service.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", service.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", service.ExcluirUsuario).Methods(http.MethodDelete)

	fmt.Println("porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
