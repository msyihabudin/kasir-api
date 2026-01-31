package handlers

import (
	"encoding/json"
	"fmt"
	model "kasir-api/models"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	categories    = make(map[string]model.Category)
	categoryMutex = &sync.Mutex{}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var c model.Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.ID = fmt.Sprintf("%d", rand.Intn(100000))

	categoryMutex.Lock()
	categories[c.ID] = c
	categoryMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	categoryMutex.Lock()
	defer categoryMutex.Unlock()

	var categoryList []model.Category
	for _, c := range categories {
		categoryList = append(categoryList, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoryList)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	categoryMutex.Lock()
	c, ok := categories[id]
	categoryMutex.Unlock()

	if !ok {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	var c model.Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.ID = id

	categoryMutex.Lock()
	if _, ok := categories[id]; !ok {
		categoryMutex.Unlock()
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	categories[id] = c
	categoryMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	categoryMutex.Lock()
	if _, ok := categories[id]; !ok {
		categoryMutex.Unlock()
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	delete(categories, id)
	categoryMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
