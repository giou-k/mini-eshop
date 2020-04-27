package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Database struct {
	*gorm.DB
}

func New() (*Database, error) {
	db, err := gorm.Open("mysql", "root:root@(localhost)/eshop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return &Database{db}, nil
}
