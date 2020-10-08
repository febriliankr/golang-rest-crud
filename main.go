package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

// GET
func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function Called: getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

// POST
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(item)

}

// DELETE
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	_deleteItemAtUID(params["uid"])
	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemAtUID(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			//Delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

// UPDATE
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) // This is like JSON.parse

	params := mux.Vars(r)

	_deleteItemAtUID(params["uid"])     // Delete the item at UID
	inventory = append(inventory, item) // append a new data

	json.NewEncoder(w).Encode(inventory)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "A fine block of cheese",
		Price: 4.99,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Wine",
		Desc:  "A fine wine",
		Price: 12.99,
	})
	inventory = append(inventory, Item{
		UID:   "2",
		Name:  "Milk",
		Desc:  "Ultra Milk",
		Price: 3.25,
	})
	handleRequests()
}
