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
	"api-kasir/middlewares"
)

type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
	ApiKey string `mapstructure:"API_KEY"`
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
		ApiKey: viper.GetString("API_KEY"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	apiKeyMiddleware := middlewares.ApiKey(config.ApiKey)
	loggingMiddleware := middlewares.Logger
	corsMiddleware := middlewares.CORS

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	// report /api/report/hari-ini
	http.HandleFunc("/api/report/hari-ini", corsMiddleware(loggingMiddleware(apiKeyMiddleware(reportHandler.GetDailyReport))))
	http.HandleFunc("/api/report", corsMiddleware(loggingMiddleware(apiKeyMiddleware(reportHandler.GetReport))))

	// checkout endpoint
	http.HandleFunc("/api/checkout", corsMiddleware(loggingMiddleware(apiKeyMiddleware(transactionHandler.HandleCheckout))))

	// /categories/{id}
	http.HandleFunc("/api/categories/", corsMiddleware(loggingMiddleware(apiKeyMiddleware(categoryHandler.HandleCategoryByID))))

	// /categories
	http.HandleFunc("/api/categories", corsMiddleware(loggingMiddleware(apiKeyMiddleware(categoryHandler.HandleCategories))))

	// HandleProductByID - GET/PUT/DELETE /api/produk/{id}
	http.HandleFunc("/api/produk/", corsMiddleware(loggingMiddleware(apiKeyMiddleware(productHandler.HandleProductByID))))
	// /api/produk
	http.HandleFunc("/api/produk", corsMiddleware(loggingMiddleware(apiKeyMiddleware(productHandler.HandleProducts))))

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
