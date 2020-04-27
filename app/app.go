package app

import (
	"github.com/giou-k/mini-eshop/database"
	"github.com/giou-k/mini-eshop/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type server struct {
	router *mux.Router
	dB     *database.Database
}

func initialMigration() (*database.Database, error) {

	db, err := database.New()

	// Migrate the schema.
	// Automatically create and update db schema.
	db.AutoMigrate(&model.Category{}, &model.Product{})

	return db, err
}

func Run() error {
	db, err := initialMigration()
	if err != nil {
		return errors.Wrap(err, "failed to connect database")
	}
	s := &server{
		dB: db,
	}

	s.routes()
	return nil
}
