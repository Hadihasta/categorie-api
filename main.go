package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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


func updateCategorieById(w http.ResponseWriter, r *http.Request){

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
	if err != nil{
		http.Error(w,"invalid request", http.StatusBadRequest)
		return
	}

	for i := range categorie{
		if categorie[i].ID == id{
			updateCategorie.ID = id
			categorie[i] = updateCategorie
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategorie)
			return
		}
	}
	http.Error(w, "categorie belum ada", http.StatusNotFound)

}


func deleteCategorieById(w http.ResponseWriter, r *http.Request){ 
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


	for i, c := range categorie{
		if c.ID == id{
			// cara menghapusnya gabungkan dari index ke i (yang mau di hapus ) dan semua sesudah index ke i (yang mau di hapus)
			categorie = append(categorie[:i], categorie[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				// best practice delete tidak return data yang di delete
				"message" : "sukses Delete",
			})
			return
		}
	}
	http.Error(w, "categorie belum ada", http.StatusNotFound)
}

func main() {

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
		}else if r.Method == "PUT"{ 
			updateCategorieById(w,r)
		}else if r.Method == "DELETE"{
			deleteCategorieById(w,r)
		}

	})

	fmt.Println("server running 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running")
	}
}
