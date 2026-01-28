package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/spf13/viper"
	"api-kasir/database"
	"api-kasir/handlers"
	"api-kasir/repositories"
	"api-kasir/services"
	"log"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// /categories/{id}
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// /categories
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)

	// HandleProductByID - GET/PUT/DELETE /api/produk/{id}
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	// /api/produk
	http.HandleFunc("/api/produk", productHandler.HandleProducts)

	// health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "server running",
		})
	})

	// root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "welcome to Api Kasir",
		})
	})

	fmt.Println("Server running di port", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
