package controllers

import (
	m "week5/models"
	"encoding/json"
	"log"
	"net/http"
)

func InsertNewTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction m.Transactions
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "ERROR : ", http.StatusBadRequest)
		log.Println("ERROR : ", err)
		return
	}

	_, err = db.Exec("INSERT INTO transactions (userID, productID, quantity) VALUES (?, ?, ?)", transaction.UserID, transaction.ProductID, transaction.Quantity)
	if err != nil {
		http.Error(w, "Failed to insert new transaction", http.StatusInternalServerError)
		log.Println("Failed to insert new transaction:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New transaction created successfully"))
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var updatedTransaction m.Transactions
	err := json.NewDecoder(r.Body).Decode(&updatedTransaction)
	if err != nil {
		http.Error(w, "ERROR : ", http.StatusBadRequest)
		log.Println("ERROR : ", err)
		return
	}

	transactionID := r.URL.Query().Get("id")
	if transactionID == "" {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		log.Println("Transaction ID is required")
		return
	}

	_, err = db.Exec("UPDATE transactions SET userID=?, productID=?, quantity=? WHERE id=?", updatedTransaction.UserID, updatedTransaction.ProductID, updatedTransaction.Quantity, transactionID)
	if err != nil {
		http.Error(w, "Failed to update transaction", http.StatusInternalServerError)
		log.Println("Failed to update transaction:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction updated successfully"))
}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("id")
	if transactionID == "" {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		log.Println("Transaction ID is required")
		return
	}

	_, err := db.Exec("DELETE FROM transactions WHERE id=?", transactionID)
	if err != nil {
		http.Error(w, "Failed to delete transaction", http.StatusInternalServerError)
		log.Println("Failed to delete transaction:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction deleted successfully"))
}

func GetDetailUserTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT t.id, t.quantity, u.id, u.name, u.age, u.address, p.id, p.name, p.price
			FROM transactions t JOIN users u ON t.userid = u.id
			JOIN products p ON t.productid = p.id`
	detailTransactionRow, err := db.Query(query)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w, "Invalid query")
		return
	}
	var detailTransaction m.DetailTransaction
	var detailTransactions []m.DetailTransaction
	for detailTransactionRow.Next() {
		if err := detailTransactionRow.Scan(
			&detailTransaction.ID, &detailTransaction.Quantity, &detailTransaction.User.ID,
			&detailTransaction.User.Name, &detailTransaction.User.Age, &detailTransaction.User.Address,
			&detailTransaction.Product.ID, &detailTransaction.Product.Name, &detailTransaction.Product.Price); err != nil {
			print(err.Error())
			sendErrorResponse(w, "Failed to scan data")
			return
		} else {
			detailTransactions = append(detailTransactions, detailTransaction)
		}
	}

	var response m.DetailTransactionsResponse
	response.Status = 200
	response.Message = "Success!"
	response.Data = detailTransactions
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, s string) {
	panic("unimplemented")
}

func isTransactionExists(productId int) bool {
	db := connect()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM transactions WHERE id=?",
		productId,
	).Scan(&count)

	if err != nil {
		log.Println("Error checking ProductID:", err)
		return false
	}

	return count > 0
}
