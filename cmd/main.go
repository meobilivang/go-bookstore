package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/meobilivang/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	// Register Routes
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	// Create a Server running on port 9010
	fmt.Printf("Starting Server at Port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
