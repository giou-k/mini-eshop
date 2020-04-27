package database

import (
	"github.com/giou-k/mini-eshop/model"
	"github.com/pkg/errors"
)

func (db *Database) ProductsGetAll() ([]*model.Product, error) {
	var categories []*model.Product
	return categories, errors.Wrap(db.Find(&categories).Error, "unable to get all categories")
}

func (db *Database) ProductCreate(product *model.Product) error {
	return errors.Wrap(db.Create(&product).Error, "unable to create product")
}

func (db *Database) ProductGetbyID(inputID string) (*model.Product, error) {
	var product model.Product

	if err := db.Where("id = ?", inputID).Find(&product).Error; err != nil {
		return nil, errors.Wrap(err, "unable to get product")
	}

	return &product, nil
}

func (db *Database) ProductDeletebyID(inputID string) error {
	var product model.Product

	if err := db.Where("id = ?", inputID).Find(&product).Error; err != nil {
		return errors.Wrap(err, "unable to find product")
	}

	return errors.Wrap(db.Delete(&product).Error, "unable to delete product")
}

func (db *Database) ProductSave(product *model.Product) error {
	return errors.Wrap(db.Save(&product).Error, "unable to save product")
}
