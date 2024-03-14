package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	m "week5/models"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age, address FROM users")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "ERROR : ", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []m.User

	for rows.Next() {
		var user m.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "ERROR : ", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	_, err = db.Exec("INSERT INTO users (name, age, address) VALUES (?, ?, ?)", user.Name, user.Age, user.Address)
	if err != nil {
		http.Error(w, "Failed to insert new user", http.StatusInternalServerError)
		log.Println("Failed to insert new user:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New user created successfully"))
}

func InsertUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")
	address := r.URL.Query().Get("address")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Input harus dalam angka :", http.StatusBadRequest)
		log.Println("Input harus dalam angka:", err)
		return
	}

	user := m.User{
		Name:    name,
		Age:     age,
		Address: address,
	}

	result := db.Create(&user)
	err = result.Error
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}
	http.Error(w, "ERROR :", http.StatusBadRequest)
	log.Println("Berhasil :", err)
	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser m.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		log.Println("User ID is required")
		return
	}

	_, err = db.Exec("UPDATE users SET name=?, age=?, address=? WHERE id=?", updatedUser.Name, updatedUser.Age, updatedUser.Address, userID)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		log.Println("Failed to update user:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func UpdateUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	var user m.User

	result := db.First(&user, &userID)
	err = result.Error
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")
	user.Name = name
	user.Address = address

	if err := db.Save(&user).Error; err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		log.Println("User ID is required")
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		log.Println("Failed to delete user:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func DeleteUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	var user m.User

	result := db.Delete(&user, &userID)
	err = result.Error
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func SelectUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		http.Error(w, "ERROR :", http.StatusBadRequest)
		log.Println("ERROR :", err)
		return
	}
	var user []m.User

	result := db.First(&user)
	if result.Error != nil {
		http.Error(w, "Gagal get User :", http.StatusBadRequest)
		log.Println("Gagal get User :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.UsersResponses
	response.Status = 200
	response.Message = "Succes"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}

func RawGorm(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		http.Error(w, "Error connect to DB :", http.StatusBadRequest)
		log.Println("Error connect to DB :", err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	var user []m.User

	result := db.Raw("SELECT name, age FROM users WHERE id = ?", &userID).Scan(&user)
	err = result.Error
	if err != nil {
		http.Error(w, "Data not found :", http.StatusBadRequest)
		log.Println("Data not found :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.UsersResponses
	response.Status = 200
	response.Message = "Succes"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}

func isUserExists(userId int) bool {
	db := connect()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id=?",
		userId,
	).Scan(&count)

	if err != nil {
		log.Println("Error checking UserID:", err)
		return false
	}

	return count > 0
}
