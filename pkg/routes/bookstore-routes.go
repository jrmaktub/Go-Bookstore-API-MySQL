package routes

import (
	"github.com/gorilla/mux"
	//import might be broken
	"github.com/jrmaktub/go-bookstore/pkg/controllers"
)

//all routes. Routes will help get the control to controller
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId", controllers.DeleteBook).Methods("DELETE")
}
