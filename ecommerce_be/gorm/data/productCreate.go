package data

import (
	"encoding/json"
	"exercise/gorm/ecommerce/models"
	"io"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func FromJSON(r io.Reader, p interface{}) error {
	newDecoder := json.NewDecoder(r)
	return newDecoder.Decode(p)
}

func AddProduct(p *models.Product) {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(p)
}
