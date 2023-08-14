package handlers

import (
	"exercise/gorm/ecommerce/data"
	"exercise/gorm/ecommerce/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET: All Products")

	lp := data.GetProducts()

	err := data.ToJSON(rw, lp)
	if err != nil {
		http.Error(rw, "JSON Conversion not successful", http.StatusInternalServerError)
	}
}

func (p *Products) GetProductDetails(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET: Product details")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unidentified id", http.StatusBadRequest)
	}

	product := data.GetProductDetails(id)
	err = data.ToJSON(rw, product)

	if err != nil {
		http.Error(rw, "Unable to marshal to JSON", http.StatusInternalServerError)
	}

}

func (p *Products) GetProductReviews(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET: Product reviews")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unidentified id", http.StatusBadRequest)
	}

	product := data.GetProductReviews(id)
	err = data.ToJSON(rw, product)

	if err != nil {
		http.Error(rw, "Unable to marshal to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST: Create Product")
	product := &models.Product{}
	err := data.FromJSON(r.Body, product)
	if err != nil {
		http.Error(rw, "Error reading product: Decode", http.StatusInternalServerError)
		p.l.Println("Error deserializing the product", err)
		return
	}
	data.AddProduct(product)
}

func (p *Products) AddReview(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST: Create Review")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unidentified id", http.StatusBadRequest)
	}
	rating := &models.Rating{}
	rating.ProductID = uint(id)
	err = data.FromJSON(r.Body, rating)
	if err != nil {
		http.Error(rw, "Error reading product: Decode", http.StatusInternalServerError)
		p.l.Println("Error deserializing the product", err)
		return
	}
	data.AddReview(rating)
}
