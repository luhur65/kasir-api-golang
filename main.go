package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Harga int `json:"harga"`
	Stok int `json:"stok"`
}

var produk = []Product{
	{ID: 1, Nama: "Mie goreng", Harga: 50000, Stok: 2},
	{ID: 2, Nama: "Mie goreng", Harga: 50000, Stok: 2},
}

func getProductByID(w http.ResponseWriter, r *http.Request){

	idrStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idrStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	
	for _, p := range produk {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "Produk tidak ada", http.StatusNotFound)

}

func updateProduct(w http.ResponseWriter, r *http.Request){

	// get id produk
	idrStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti int
	id, err := strconv.Atoi(idrStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	
	// get data dari req
	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// loop produk
	for i := range produk {
		if produk[i].ID == id {
			updateProduct.ID = id
			produk[i] = updateProduct
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "Produk tidak ada", http.StatusNotFound)

}

func deleteProduct(w http.ResponseWriter, r *http.Request){

	// get id produk
	idrStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti id int
	id, err := strconv.Atoi(idrStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	// loop produk
	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string {
				"status": "ok",
				"message": "berhasil hapus",
			})
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "Produk tidak ada", http.StatusNotFound)

}

func main()  {

	// GET produk detail localhost:8080/api/produk/{id}
	// PUT produk detail localhost:8080/api/produk/{id}
	// DELETE produk detail localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}

	})


	// GET
	// POST
	// localhost:8080/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Product
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}
		// w.Write([]byte("OK"))
	}) 

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string {
			"status": "ok",
			"message": "server running",
		})
		// w.Write([]byte("OK"))
	}) 

	fmt.Println("Server running di port 8080")

	// cara buat server di port 8080 
	err := http.ListenAndServe(":8080", nil)

	// cek jika server gagal terhubung
	if err != nil {
		fmt.Println("gagal running server")
	}
	
}