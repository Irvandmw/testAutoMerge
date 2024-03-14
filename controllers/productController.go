package controllers

import (
	m "week5/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ERROR : ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []m.Products

	for rows.Next() {
		var product m.Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "ERROR : ", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func InsertNewProduct(w http.ResponseWriter, r *http.Request) {
	var product m.Products
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	_, err = db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		http.Error(w, "Failed to insert new product", http.StatusInternalServerError)
		log.Println("Failed to insert new product:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New product created successfully"))
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct m.Products
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "ERROR ", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		log.Println("Product ID is required")
		return
	}

	_, err = db.Exec("UPDATE products SET name=?, price=? WHERE id=?", updatedProduct.Name, updatedProduct.Price, productID)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		log.Println("Failed to update product:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		log.Println("Product ID is required")
		return
	}

	_, err := db.Exec("DELETE FROM products WHERE id=?", productID)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		log.Println("Failed to delete product:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}

// Cek apakah product dengan ID tertentu sudah ada atau belum
func isProductExists(productId int) bool {
	db := connect()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products WHERE id=?",
		productId,
	).Scan(&count)

	if err != nil {
		log.Println("Error checking ProductID:", err)
		return false
	}

	return count > 0
}