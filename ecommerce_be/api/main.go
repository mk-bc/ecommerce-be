package main

import (
	"exericse/ecommerce/api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// custom logger for handlers
	l := log.New(os.Stdout, "ecommerce-api: ", log.LstdFlags)

	// product handler
	productHandler := handlers.NewProducts(l)

	smux := mux.NewRouter()

	// handlers for get method
	getRouter := smux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/products", productHandler.GetProducts)
	getRouter.HandleFunc("/api/products/{id:[0-9]+}", productHandler.GetProductDetails)
	getRouter.HandleFunc("/api/products/{id:[0-9]+}/reviews", productHandler.GetProductReviews)

	// handlers for post method
	postRouter := smux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/products/create", productHandler.AddProduct)
	postRouter.HandleFunc("/api/products/{id:[0-9]+}/reviews/create", productHandler.AddReview)

	// setting up the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: smux,
	}

	// running the server in a go routine
	// go func() {
	// 	l.Println("Starting server on port:8080")
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()

	server.ListenAndServe()

}
