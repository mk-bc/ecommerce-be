package data

import (
	"encoding/json"
	"exercise/gorm/ecommerce/models"
	"io"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ToJSON(w io.Writer, p interface{}) error {
	newEncoder := json.NewEncoder(w)
	// fmt.Println(p)
	return newEncoder.Encode(p)
}

type ProductView struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Image       string `json:"image"`
}
type Products []*ProductView

func GetProducts() Products {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	lp := []*ProductView{}
	// db.Debug().Model(&models.Product{}).Select([]string{"id", "name", "description", "category", "quantity", "price", "image"}).Find(&lp)
	rows, err := db.Debug().Model(&models.Product{}).Select([]string{"id", "name", "description", "category", "quantity", "price", "image"}).Rows()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var p ProductView
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Quantity, &p.Price, &p.Image)
		lp = append(lp, &p)
	}

	// for _, s := range lp {
	// 	fmt.Println(s)
	// }
	return lp
}

type ProductDetailView struct {
	// productView ProductView
	ID          uint
	Name        string
	Description string
	Category    string
	Quantity    uint
	Price       uint
	Image       string
	variants    []*models.Variant
}

func GetProductDetails(id int) models.Product {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	product := models.Product{}

	db.Debug().Model(&models.Product{}).Preload("Variants").Where("id = ?", id).First(&product)

	result := ProductDetailView{}
	result.variants = product.Variants
	// result.productView = ProductView{
	// 	ID:          product.ID,
	// 	Name:        product.Name,
	// 	Description: product.Description,
	// 	Category:    product.Category,
	// 	Quantity:    product.Quantity,
	// 	Price:       product.Price,
	// 	Image:       product.Image,
	// }
	result.ID = product.ID
	result.Name = product.Name
	result.Description = product.Description
	result.Category = product.Category
	result.Quantity = product.Quantity
	result.Price = product.Price
	result.Image = product.Image

	return product

}

func GetProductReviews(id int) models.Product {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	product := models.Product{}

	db.Debug().Model(&models.Product{}).Preload("Ratings").Where("id = ?", id).First(&product)

	return product
}
