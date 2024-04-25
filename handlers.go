package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item

	json.NewDecoder(r.Body).Decode(&newItem) // Tre sa trimit referinta.
	newItem.ID = uuid.New().String()
	items = append(items, newItem)

	json.NewEncoder(w).Encode(newItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updatedItem Item

	json.NewDecoder(r.Body).Decode(&updatedItem)
	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i].Name = updatedItem.Name

			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var toDelete Item

	json.NewDecoder(r.Body).Decode(&toDelete)
	for i, item := range items {
		if item.ID == toDelete.ID {
			items = append(items[:i], items[i+1:]...) // (...) Trateze slice-ul items[i+1:] ca o lista de argumente separate.

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
