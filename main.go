package main

import (
	"categories-api/database"
	"categories-api/handlers"
	"categories-api/repositories"
	"categories-api/services"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

type Categorie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categorie = []Categorie{
	{ID: 1,
		Name:        "minuman",
		Description: "semua hal yang minuman"},
	{ID: 2,
		Name:        "makanan",
		Description: "semua hal yang makanan"},
}

func getCategorieById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Categorie ID", http.StatusBadRequest)
		return
	}

	for _, p := range categorie {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "categorie belum ada", http.StatusNotFound)

}

func updateCategorieById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// ambil ID
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateCategorie Categorie
	err = json.NewDecoder(r.Body).Decode(&updateCategorie)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	for i := range categorie {
		if categorie[i].ID == id {
			updateCategorie.ID = id
			categorie[i] = updateCategorie

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategorie)
			return
		}
	}
	http.Error(w, "categorie belum ada", http.StatusNotFound)

}

func deleteCategorieById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/categorie/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, c := range categorie {
		if c.ID == id {
			// cara menghapusnya gabungkan dari index ke i (yang mau di hapus ) dan semua sesudah index ke i (yang mau di hapus)
			categorie = append(categorie[:i], categorie[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				// best practice delete tidak return data yang di delete
				"message": "sukses Delete",
			})
			return
		}
	}
	http.Error(w, "categorie belum ada", http.StatusNotFound)
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
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/product", productHandler.HandleProducts)

	// localhost 8080/health
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})

	// localhost 8080/categorie
	http.HandleFunc("/api/categorie", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categorie)
		} else if r.Method == "POST" {
			var categorieBaru Categorie
			err := json.NewDecoder(r.Body).Decode(&categorieBaru)
			if err != nil {
				http.Error(w, "invalid request...", http.StatusBadRequest)
				return
			}
			categorieBaru.ID = len(categorie) + 1
			categorie = append(categorie, categorieBaru)

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categorieBaru)
		}
	})

	http.HandleFunc("/api/categorie/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategorieById(w, r)
		} else if r.Method == "PUT" {
			updateCategorieById(w, r)
		} else if r.Method == "DELETE" {
			deleteCategorieById(w, r)
		}

	})

		// format for listen and serve :(port)
	addr := ":" + config.Port

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
