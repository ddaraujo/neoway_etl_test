package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Get all valid items on database
func getAllValidItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items")
	items, err := dbInstance.GetAllValidItems()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all invalid (rejected) items on database
func getAllInvalidItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all invalid items")
	items, err := dbInstance.GetAllInvalidItems()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all items on database by cpf
func getAllItemsByCpf(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items by CPF")
	vars := mux.Vars(r)
	cpf := vars["cpf"]
	items, err := dbInstance.GetAllItemsByCpf(cpf)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all items on database by cpf
func getAllItemsLastSaleCnpj(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items by last sale CNPJ")
	vars := mux.Vars(r)
	cnpj := vars["cnpj"]
	items, err := dbInstance.GetAllItemsLastSale(cnpj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all items on database by cpf
func getAllItemsFrequentSaleCnpj(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items by frequent sale CNPJ")
	vars := mux.Vars(r)
	cnpj := vars["cnpj"]
	items, err := dbInstance.GetAllItemsFrequentSale(cnpj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all valid records count
func getValidRecordsCount(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items count")
	items, err := dbInstance.GetValidRecordsCount()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Get all invalid records count
func getInvalidRecordsCount(w http.ResponseWriter, r *http.Request) {
	log.Println("Obtaining all valid items count")
	items, err := dbInstance.GetInvalidRecordsCount()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(items)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

// Delete all valid items
/*
func deleteValidItems(w http.ResponseWriter, r *http.Request) {
	err := dbInstance.DeleteValidItems()
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}*/
