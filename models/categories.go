package models

type Categories struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// var categories = []Category{
// 	{ID: 1, Name: "Makanan", Description: "Produk makanan"},
// 	{ID: 2, Name: "Minuman", Description: "Produk minuman"},
// }


// func getCategories(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(categories)
// }


// func getCategoryByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, c := range categories {
// 		if c.ID == id {
// 			json.NewEncoder(w).Encode(c)
// 			return
// 		}
// 	}

// 	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
// }


// func createCategory(w http.ResponseWriter, r *http.Request) {
// 	var newCategory Category
// 	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	newCategory.ID = len(categories) + 1
// 	categories = append(categories, newCategory)

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newCategory)
// }


// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	var updated Category
// 	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	for i := range categories {
// 		if categories[i].ID == id {
// 			updated.ID = id
// 			categories[i] = updated
// 			json.NewEncoder(w).Encode(updated)
// 			return
// 		}
// 	}

// 	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
// }


// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	for i, c := range categories {
// 		if c.ID == id {
// 			categories = append(categories[:i], categories[i+1:]...)
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"status":  "ok",
// 				"message": "category berhasil dihapus",
// 			})
// 			return
// 		}
// 	}

// 	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
// }
