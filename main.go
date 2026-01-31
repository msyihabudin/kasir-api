package main

import (
	"context"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/middleware"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DbConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat("configs/.env"); err == nil {
		viper.SetConfigFile("configs/.env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DbConn: viper.GetString("DB_CONN"),
	}

	// Initialize Database
	db, err := database.InitDB(config.DbConn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close(context.Background())

	// Initialize Repositories
	productRepo := repositories.NewProductRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// Initialize Services
	productService := services.NewProductService(productRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	// Initialize Handlers
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup Routes
	SetupRoutes(productHandler, categoryHandler)

	// Wrap all handlers with TraceIDMiddleware
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.TraceIDMiddleware(http.DefaultServeMux).ServeHTTP(w, r)
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on port ", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
