package main

import (
	"./helpers"
	_ "./helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/get", apiUsersGet).Methods("POST")
	r.HandleFunc("/users/get/{id}", apiUsersGetById).Methods("POST")

	r.HandleFunc("/orders/get", apiOrdersGet).Methods("POST")
	r.HandleFunc("/orders/get/{id}", apiOrdersGetById).Methods("GET")
	r.HandleFunc("/orders/create", apiOrdersCreate).Methods("GET")

	r.HandleFunc("/pubs/getlist", apiPubsGetList).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func apiPubsGetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	pubs, error := helpers.GetPubsList()
	if !error {
		json.NewEncoder(w).Encode(pubs)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func apiOrdersCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	order, error := helpers.CreateOrder()
	if !error {
		json.NewEncoder(w).Encode(order)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func apiUsersGetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(vars["id"], 16, 32)
	user, error := helpers.GetUserById(userID)
	if !error {
		json.NewEncoder(w).Encode(user)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func apiOrdersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	order, error := helpers.GetOrders()
	if !error {
		json.NewEncoder(w).Encode(order)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func apiUsersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	users, error := helpers.GetUsers()
	if !error {
		json.NewEncoder(w).Encode(users)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func apiOrdersGetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	vars := mux.Vars(r)
	orderID, _ := strconv.ParseInt(vars["id"], 16, 32)
	orders, error := helpers.GetOrderById(orderID)
	if !error {
		json.NewEncoder(w).Encode(orders)
	} else {
		message := Message{
			Code:    500,
			Message: "Problem whit database",
		}
		json.NewEncoder(w).Encode(message)
	}

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{
		Code:    404,
		Message: "Oops..",
	})
}
