package data

import (
	"exercise/gorm/ecommerce/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func AddReview(r *models.Rating) {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(r)
}
