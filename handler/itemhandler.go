package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ohlmeier/todo-api/database"
	"github.com/ohlmeier/todo-api/todo"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item todo.Item
	json.NewDecoder(r.Body).Decode(&item)
	database.Instance.Create(&item)
	json.NewEncoder(w).Encode(item)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if !listExists(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var item todo.Item
	database.Instance.First(&item, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []todo.Item
	database.Instance.Find(&items)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if !listExists(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var item todo.Item
	database.Instance.First(&item, id)
	json.NewDecoder(r.Body).Decode(&item)
	database.Instance.Save(&item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if !listExists(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var item todo.Item
	database.Instance.Delete(&item, id)
	json.NewEncoder(w).Encode("Item Deleted Successfully")
}

func itemExists(itemId string) bool {
	var item todo.Item
	database.Instance.First(&item, itemId)
	return item.ID != 0
}
