package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ohlmeier/todo-api/database"
	"github.com/ohlmeier/todo-api/handler"
)

func main() {

	connectionString := os.Getenv("DB_CONNECTION")
	database.Connect(connectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(false)

	registerRoutes(router)

	port := os.Getenv("PORT")
	log.Printf("Starting Server on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/todo/lists", handler.GetLists).Methods("GET")
	router.HandleFunc("/todo/lists/{id}", handler.GetListById).Methods("GET")
	router.HandleFunc("/todo/lists", handler.CreateList).Methods("POST")
	router.HandleFunc("/todo/lists/{id}", handler.UpdateList).Methods("PUT")
	router.HandleFunc("/todo/lists/{id}", handler.DeleteList).Methods("DELETE")
	router.HandleFunc("/todo/lists/{id}/items", handler.GetItems).Methods("GET")
	router.HandleFunc("/todo/lists/{id}/items/{id}", handler.GetItemById).Methods("GET")
	router.HandleFunc("/todo/lists/{id}/items", handler.CreateItem).Methods("POST")
	router.HandleFunc("/todo/lists/{id}/items/{id}", handler.UpdateItem).Methods("PUT")
	router.HandleFunc("/todo/lists/{id}/items/{id}", handler.DeleteItem).Methods("DELETE")
}
