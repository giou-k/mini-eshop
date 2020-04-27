package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *server) routes() {
	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc("/categories", s.categoriesGetAll).Methods("GET")
	s.router.HandleFunc("/category/{id}", s.categoryGetbyID).Methods("GET")
	s.router.HandleFunc("/category/{id}", s.categoryDeletebyID).Methods("DELETE")
	s.router.HandleFunc("/category/{id}", s.categoryUpdatebyID).Methods("PATCH")
	s.router.HandleFunc("/category", s.categoryCreate).Methods("POST")
	//

	s.router.HandleFunc("/products", s.productGetAll).Methods("GET")
	s.router.HandleFunc("/product/{id}", s.productGetByID).Methods("GET")
	s.router.HandleFunc("/product/{id}", s.productDeleteByID).Methods("DELETE")
	s.router.HandleFunc("/product/{id}", s.productUpdateByID).Methods("PATCH")
	s.router.HandleFunc("/product", s.productCreate).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", s.router))
}
