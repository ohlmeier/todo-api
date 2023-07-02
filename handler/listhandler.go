package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohlmeier/todo-api/database"
	"github.com/ohlmeier/todo-api/todo"
)

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var list todo.List
	json.NewDecoder(r.Body).Decode(&list)
	database.Instance.Create(&list)
	json.NewEncoder(w).Encode(list)
}

func GetListtById(w http.ResponseWriter, r *http.Request) {
	listId := mux.Vars(r)["id"]
	if !listExists(listId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var list todo.List
	database.Instance.First(&list, listId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func GetLists(w http.ResponseWriter, r *http.Request) {
	var lists []todo.List
	database.Instance.Find(&lists)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	listId := mux.Vars(r)["id"]
	if !listExists(listId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var list todo.List
	database.Instance.First(&list, listId)
	json.NewDecoder(r.Body).Decode(&list)
	database.Instance.Save(&list)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listId := mux.Vars(r)["id"]
	if !listExists(listId) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var list todo.List
	database.Instance.Delete(&list, listId)
	json.NewEncoder(w).Encode("List Deleted Successfully")
}

func listExists(listId string) bool {
	var list todo.List
	database.Instance.First(&list, listId)
	return list.ID != 0
}
