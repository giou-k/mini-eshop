package database

import (
	"github.com/giou-k/mini-eshop/model"
	"github.com/pkg/errors"
)

func (db *Database) CategoriesGetAll() ([]*model.Category, error) {
	var categories []*model.Category
	return categories, errors.Wrap(db.Find(&categories).Error, "unable to get all categories")
}

func (db *Database) CategoryCreate(category *model.Category) error {
	return errors.Wrap(db.Create(&category).Error, "unable to create category")
}

func (db *Database) CategoryGetbyID(inputID string) (*model.Category, error) {
	var category model.Category

	if err := db.Where("id = ?", inputID).Find(&category).Error; err != nil {
		return nil, errors.Wrap(err, "unable to get category")
	}

	return &category, nil
}

func (db *Database) CategoryDeletebyID(inputID string) error {
	var category model.Category

	if err := db.Where("id = ?", inputID).Find(&category).Error; err != nil {
		return errors.Wrap(err, "unable to find category")
	}

	return errors.Wrap(db.Delete(&category).Error, "unable to delete category")
}

func (db *Database) CategorySave(category *model.Category) error {
	return errors.Wrap(db.Save(&category).Error, "unable to save category")
}
