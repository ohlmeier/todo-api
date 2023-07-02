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
	router.HandleFunc("/api/todolists", handler.GetLists).Methods("GET")
	router.HandleFunc("/api/todolists/{id}", handler.GetListtById).Methods("GET")
	router.HandleFunc("/api/todolists", handler.CreateList).Methods("POST")
	router.HandleFunc("/api/todolists/{id}", handler.UpdateList).Methods("PUT")
	router.HandleFunc("/api/todolists/{id}", handler.DeleteList).Methods("DELETE")
}
