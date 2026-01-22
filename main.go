package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	// /api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getProductByID(w, r)
		case "PUT":
			updateProduct(w, r)
		case "DELETE":
			deleteProduct(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// /api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				getProducts(w, r)
			case "POST":
				createProduct(w, r)
			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "server running",
		})
	})

	fmt.Println("Server running di port 8080")
	http.ListenAndServe(":8080", nil)
}
