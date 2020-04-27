package app

import (
	"encoding/json"
	"fmt"
	"github.com/giou-k/mini-eshop/model"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// TODO delete
//comment := Comment{Content: "Good post!", Author: "Joe"}
//Db.Model(&post).Association("Comments").Append(comment)

func (s *server) productGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all Products...")

	categories, err := s.dB.CategoriesGetAll()
	if err != nil {
		logrus.Error(err)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (s *server) productCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a New Product...")
	var product model.Product

	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	return err
	//}

	json.Unmarshal(reqBody, &product)

	// Create product in mysql.
	err := s.dB.ProductCreate(&product)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "New Product Successfully Created")
}

func (s *server) productGetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Query a Product...")

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	product, err := s.dB.ProductGetbyID(inputID)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "Successfully queried Product: ")
	json.NewEncoder(w).Encode(product)
}

func (s *server) productDeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a Product...")

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	err := s.dB.ProductDeletebyID(inputID)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "Successfully Deleted Product")
}

func (s *server) productUpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a Product...")
	var updatedProduct model.Product

	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)

	// product from request
	json.Unmarshal(reqBody, &updatedProduct)

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	// TODO transaction
	//db.Model(&product).Updates(updatedProduct)
	product, err := s.dB.ProductGetbyID(inputID)
	if err != nil {
		logrus.Error(err)
	}


	if done := productUpdateDiffs(updatedProduct, product); done {
		// After updates where done save the new product.
		s.dB.ProductSave(product)

		fmt.Fprintf(w, "Successfully Updated Product")
		return
	}

	fmt.Fprintf(w, "Didn't get any updated input.")

}

func productUpdateDiffs(updatedProduct model.Product, product *model.Product) bool {

	// The user can delete only the fields that don't get used by our db. If a field has a new value and it's not empty,
	// then delete it.
	updated := false
	if product.Description != updatedProduct.Description && updatedProduct.Description != "" {
		product.Description = updatedProduct.Description
		updated = true
	}
	if product.Title != updatedProduct.Title && updatedProduct.Title != "" {
		product.Title = updatedProduct.Title
		updated = true
	}
	if product.ImageURL != updatedProduct.ImageURL && updatedProduct.ImageURL != "" {
		product.ImageURL = updatedProduct.ImageURL
		updated = true
	}

	return updated
}
