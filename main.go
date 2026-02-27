package main

import (
	"categories-api/database"
	"categories-api/handlers"
	"categories-api/middlewares"
	"categories-api/repositories"
	"categories-api/services"
	"encoding/json"

	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
	APiKey string `mapstructure:"API_KEY"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
		APiKey: viper.GetString("API_KEY"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	apiKeyMiddleware := middlewares.APIKEY(config.APiKey)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)
	// using middleware
	http.HandleFunc("/api/product/", apiKeyMiddleware(productHandler.HandleProductByID))
	http.HandleFunc("/api/category", categoryHandler.HandleCategorys)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// using middleware
	http.HandleFunc("/api/checkout", apiKeyMiddleware(transactionHandler.Checkout))

	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	http.HandleFunc("/api/report", reportHandler.GetReport)
	http.HandleFunc("/api/report/hari-ini", reportHandler.GetTodayReport)

	// localhost 8080/health
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})
	port := config.Port
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Println("server running on", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = db
}
