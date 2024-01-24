package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"postgresql/handler_test/models"
	"postgresql/handler_test/storage"
	"strconv"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/user/create", CreateUser)

	http.HandleFunc("/user/all", GetAllUsers)

	http.HandleFunc("/user/update", UpdateUsers)

	http.HandleFunc("/user/get", GetUserrById)

	http.HandleFunc("/user/delete", DeleteUser)

	log.Println("Server is running ....")
	
	err := http.ListenAndServe("localhost:5151", nil)
	if err != nil {
		log.Print("error while runnig server")

	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *models.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("while unmarshalling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uuid.NewString()

	user.Id = id

	respUser, err := storage.CreteUser(user)

	if err != nil {
		log.Println("error while creting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while creting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content_Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(page)

	if err != nil {
		log.Println("erro while converting page, is not integer", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")

	intLimit, err := strconv.Atoi(limit)

	if err != nil {
		log.Println("error while converting limit, is not integer", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := storage.GetAll(intPage, intLimit)

	if err != nil {
		log.Println("error while getting all users, smth went wrong", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBodyusers, err := json.Marshal(user)

	if err != nil {
		log.Println("error while creting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyusers)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	newName := r.URL.Query().Get("newName")
	newLastName := r.URL.Query().Get("newLastName")

	user, err := storage.UpdateUser(userId, newName, newLastName)
	if err != nil {
		log.Println("error while getting all users, smth went wrong", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resCreateUser, err := json.Marshal(user)
	if err != nil {
		log.Println("error while creting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resCreateUser)
}

func GetUserrById(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")

	user, err := storage.GetUser(userId)
	if err != nil {
		log.Println("error while getting user, smth went wrong", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resCreateUser, err := json.Marshal(user)
	if err != nil {
		log.Println("error while creting user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resCreateUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")

	storage.DeleteUser(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("deletted user"))
}
