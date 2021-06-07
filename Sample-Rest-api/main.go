package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
)

func calledGETMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"method:" : "GET Operation","message":"Hello-World"}`))
}

func calledPOSTMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"method:" : "POST Operation","message":"Hello-World"}`))
}
func calledPUTMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"method:" : "PUT Operation","message":"Hello-World"}`))
}

func calledDeleteMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"method:" : "DELETE Operation","message":"Hello-World"}`))
}

func calledNoMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	bookId := pathParams["bookID"]
	log.Printf(bookId)

}

func main() {
	getBookCollections := ReadCsv()
	fmt.Println(getBookCollections)
	route := mux.NewRouter()
	api := route.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", calledGETMethod).Methods(http.MethodGet)
	api.HandleFunc("", calledPOSTMethod).Methods(http.MethodPost)
	api.HandleFunc("", calledPUTMethod).Methods(http.MethodPut)
	api.HandleFunc("", calledDeleteMethod).Methods(http.MethodDelete)
	api.HandleFunc("", calledNoMethod)
	api.HandleFunc("/books/{bookID}", getBookByID).Methods(http.MethodGet)
	fmt.Println("Starting Server and  Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}
