package main

import (
	handler "kasir-api/handlers"
	"net/http"
)

// SetupRoutes registers all HTTP routes for the API
func SetupRoutes(productHandler *handler.ProductHandler, categoryHandler *handler.CategoryHandler) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"OK","message":"API Running"}`))
	})

	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)
}
