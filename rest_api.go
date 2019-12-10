//package main
//
//import (
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"log"
//	"math/rand"
//	"net/http"
//	"strconv"
//)
//
//type Book struct {
//	ID      string  `json:"id"`
//	Title   string  `json:"title"`
//	Author  *Author `json:"author"`
//	Aliases []Alias `json:"aliases"`
//}
//type Alias struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}
//type Author struct {
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//}
//
//var books []Book
//
//func getBooks(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "Application/json")
//	json.NewEncoder(w).Encode(books)
//}
//func getBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for _, item := range books {
//		if item.ID == params["id"] {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Book{})
//}
//func createBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var book Book
//	_ = json.NewDecoder(r.Body).Decode(&book)
//	book.ID = strconv.Itoa(rand.Intn(1000000))
//	books = append(books, book)
//	json.NewEncoder(w).Encode(book)
//}
//func main() {
//	r := mux.NewRouter()
//	var aliases, aliases2 []Alias
//	aliases = append(aliases, Alias{
//		ID:   "1",
//		Name: "Test 1",
//	}, Alias{
//		ID:   "2",
//		Name: "Test 2",
//	}, Alias{
//		ID:   "3",
//		Name: "Test 3",
//	})
//	aliases2 = append(aliases2, Alias{
//		ID:   "5",
//		Name: "Test 5",
//	})
//	books = append(books,
//		Book{ID: "1", Title: "Война и Мир",
//			Author:  &Author{FirstName: "Лев", LastName: "Толстой"},
//			Aliases: aliases,
//		})
//	books = append(books,
//		Book{ID: "2",
//			Title:   "Преступление и наказание",
//			Author:  &Author{FirstName: "Фёдор", LastName: "Достоевский"},
//			Aliases: aliases2,
//		})
//	r.HandleFunc("/books", getBooks).Methods("GET")
//	r.HandleFunc("/books/{id}", getBook).Methods("GET")
//	r.HandleFunc("/books", createBook).Methods("POST")
//	log.Fatal(http.ListenAndServe(":8000", r))
//}

package main

import (
	"./helpers"
	_ "./helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users.get", apiUsersGet).Methods("GET")
	r.HandleFunc("/orders.get/{id}", apiUsersGet).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{
		Code:    404,
		Message: "Oops..",
	})
}

func apiUsersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(helpers.GetUsers())
}

func apiOrdersGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.ParseInt(vars["id"])
	json.NewEncoder(w).Encode(helpers.GetOrderById(orderID))
}
