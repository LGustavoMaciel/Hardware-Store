package api

import (
	"flag"
	"fmt"
	"project/api/database"
)

var (
	port = flag.Int("p", 5000, "set port")
)

func Run() {
	flag.Parse()

	db := database.Connect()

	if db != nil {
		defer db.Close()
	}

	

	fmt.Println("Api running...")
}