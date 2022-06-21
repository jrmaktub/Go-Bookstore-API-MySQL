package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jrmaktub/go-bookstore/pkg/routes"
)

//main file tells bookstore.routes where the routes are
func main() {
	r := mux.NewRouter()
	//passing the r Router to the registerbook func in routes
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost: 9010", r))
}
