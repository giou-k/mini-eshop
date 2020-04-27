package app

import (
	"encoding/json"
	"fmt"
	"github.com/giou-k/mini-eshop/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (s *server) categoriesGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all Categories...")

	categories, err := s.dB.CategoriesGetAll()
	if err != nil {
		logrus.Error(err)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (s *server) categoryCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a New Category...")
	var category model.Category

	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	return err
	//}

	json.Unmarshal(reqBody, &category)

	// Create category in mysql.
	err := s.dB.CategoryCreate(&category)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "New Category Successfully Created")
}

func (s *server) categoryGetbyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Query a Category...")

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	category, err := s.dB.CategoryGetbyID(inputID)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "Successfully queried Category: ")
	json.NewEncoder(w).Encode(category)
}

func (s *server) categoryDeletebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a Category...")

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	err := s.dB.CategoryDeletebyID(inputID)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Fprintf(w, "Successfully Deleted Category")
}

func (s *server) categoryUpdatebyID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a Category...")
	var updatedCategory model.Category

	defer r.Body.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)

	// category from request
	json.Unmarshal(reqBody, &updatedCategory)

	// get the id from the URI
	vars := mux.Vars(r)
	inputID := vars["id"]

	// TODO transaction
	category, err := s.dB.CategoryGetbyID(inputID)
	if err != nil {
		logrus.Error(err)
	}
	//db.Model(&category).Updates(updatedCategory)

	if done := categoryUpdateDiffs(updatedCategory, category); done {
		// After updates where done save the new category.
		s.dB.CategorySave(category)

		fmt.Fprintf(w, "Successfully Updated Category")
		return
	}

	fmt.Fprintf(w, "Didn't get any updated input.")
}

func categoryUpdateDiffs(updatedCategory model.Category, category *model.Category) bool {

	// The user can delete only the fields that don't get used by our db. If a field has a new value and it's not empty,
	// then delete it.
	updated := false
	if category.Position != updatedCategory.Position && updatedCategory.Position != 0 {
		category.Position = updatedCategory.Position
		updated = true
	}
	if category.Title != updatedCategory.Title && updatedCategory.Title != "" {
		category.Title = updatedCategory.Title
		updated = true
	}
	if category.ImageURL != updatedCategory.ImageURL && updatedCategory.ImageURL != "" {
		category.ImageURL = updatedCategory.ImageURL
		updated = true
	}

	return updated
}
