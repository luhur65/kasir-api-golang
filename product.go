package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Product{
	{ID: 1, Nama: "Mie goreng", Harga: 10000, Stok: 2},
	{ID: 2, Nama: "Nasi goreng", Harga: 20000, Stok: 2},
}

// ================= HANDLER =================

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
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

	http.Error(w, "Produk tidak ada", http.StatusNotFound)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var produkBaru Product
	if err := json.NewDecoder(r.Body).Decode(&produkBaru); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produkBaru)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	var update Product
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			update.ID = id
			produk[i] = update
			json.NewEncoder(w).Encode(update)
			return
		}
	}

	http.Error(w, "Produk tidak ada", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "ok",
				"message": "berhasil hapus",
			})
			return
		}
	}

	http.Error(w, "Produk tidak ada", http.StatusNotFound)
}
