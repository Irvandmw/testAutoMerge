package main

import (
	"week5/controllers"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	controllers.InitializeDB(db)
	router := mux.NewRouter()

	// USER
	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	// router.HandleFunc("/users", controllers.InsertNewUser).Methods("POST")
	// router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users", controllers.Login).Methods("POST")

	// // PRODUCT
	// router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	// router.HandleFunc("/products", controllers.InsertNewProduct).Methods("POST")
	// router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products", controllers.DeleteProduct).Methods("DELETE")

	// // TRANSACTION
	// router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactions", controllers.InsertNewTransaction).Methods("POST")
	// router.HandleFunc("/transactions", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions", controllers.DeleteTransaction).Methods("DELETE")
	// router.HandleFunc("/transactions", controllers.GetDetailUserTransaction).Methods("GET")

	

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
