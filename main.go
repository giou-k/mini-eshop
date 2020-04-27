package main

import (
	"fmt"
	"github.com/giou-k/mini-eshop/app"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	os.Exit(1)
}
