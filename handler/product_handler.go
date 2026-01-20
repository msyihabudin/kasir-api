package handler

import (
	"encoding/json"
	"fmt"
	"kasir-api/model"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	products = make(map[string]model.Product)
	mutex    = &sync.Mutex{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.ID = fmt.Sprintf("%d", rand.Intn(100000))

	mutex.Lock()
	products[p.ID] = p
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	var productList []model.Product
	for _, p := range products {
		productList = append(productList, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	p, ok := products[id]
	mutex.Unlock()

	if !ok {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	var p model.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = id

	mutex.Lock()
	if _, ok := products[id]; !ok {
		mutex.Unlock()
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	products[id] = p
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if _, ok := products[id]; !ok {
		mutex.Unlock()
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	delete(products, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
