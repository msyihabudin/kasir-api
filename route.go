package main

import (
	handler "kasir-api/handlers"
	"net/http"
)

// SetupRoutes registers all HTTP routes for the API
func SetupRoutes(productHandler *handler.ProductHandler) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"OK","message":"API Running"}`))
	})

	http.HandleFunc("/api/products", productHandler.HandleProducts)

	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	// http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handler.GetCategories(w, r)
	// 	case http.MethodPost:
	// 		handler.CreateCategory(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })

	// http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		handler.GetCategory(w, r)
	// 	case http.MethodPut:
	// 		handler.UpdateCategory(w, r)
	// 	case http.MethodDelete:
	// 		handler.DeleteCategory(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })
}
