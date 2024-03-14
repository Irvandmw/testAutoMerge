package controllers

import (
	m "week5/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
	db = database
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "Failed to parse request body")
		return
	}

	platform := r.Header.Get("platform")
	if platform == "" {
		sendErrorResponse(w, "Missing platform header")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user m.User

	errQuery := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?",
		email,
		password,
	).Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType)

	var response m.UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success login from " + platform
		response.Data = user
	} else {
		response.Status = 400
		response.Message = "Login failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, userID, productID, quantity FROM transactions")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ERROR : ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []m.Transactions

	for rows.Next() {
		var transaction m.Transactions
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "ERROR : ", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// * USERS


// * PRODUCTS


// * TRANSACTIONS
