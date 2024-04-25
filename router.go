package main

import (
	"github.com/gorilla/mux"
)

func NewServer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/item", AddItem).Methods("POST")
	router.HandleFunc("/item", UpdateItem).Methods("PUT")
	router.HandleFunc("/item", DeleteItem).Methods("DELETE")

	return router
}
