package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	seedDB(db)

	// manual insertion of data
	// inserting products
	for _, product := range products {
		db.Debug().Create(&product)
	}

}

func seedDB(db *gorm.DB) {
	db.DropTableIfExists(&Rating{})
	// db.CreateTable(&Rating{})
	db.DropTableIfExists(&Variant{})
	// db.CreateTable(&Variant{})
	db.DropTableIfExists(&Product{})
	// db.CreateTable(&Product{})
	db.AutoMigrate(&Product{}, &Rating{}, &Variant{})
}

func (p Product) TableName() string {
	return "products"
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	Quantity    uint
	Price       uint
	Image       string
	Ratings     []*Rating
	Variants    []*Variant
}

type Variant struct {
	ID        uint
	Color     string
	Image     string
	ProductID uint
}

type Rating struct {
	gorm.Model
	Name      string
	Review    string
	Rating    uint
	ProductID uint
}

var products []*Product = []*Product{
	{
		Name:        "Bottle",
		Description: "No leakage",
		Category:    "Kitchen",
		Quantity:    15,
		Price:       320,
		Image:       "https://image.shutterstock.com/image-photo/stylish-stainless-thermo-bottles-on-260nw-1914561409.jpg",
		Variants: []*Variant{
			{
				Color: "Grey",
				Image: "https://images.unsplash.com/photo-1602143407151-7111542de6e8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2F0ZXIlMjBib3R0bGV8ZW58MHx8MHx8&w=1000&q=80",
			},
			{
				Color: "Blue",
				Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCzbSXJvxUSmMN9JlQGMy3hVmv_GBkrFl1YQ&usqp=CAU",
			},
		},
	},
	{
		Name:        "Dell Laptop | i7 | 16inch | 8GB RAM | 256 GB Storage | 500MG GRAPHIC card",
		Description: "Powered by intel processor and blazing fast laptop coming in 3 variants",
		Category:    "Electronics",
		Quantity:    10,
		Image:       "https://s.yimg.com/uu/api/res/1.2/NVx3iijOkY6GsTUJNkGIAg--~B/aD0xMjAwO3c9MjAwMDthcHBpZD15dGFjaHlvbg--/https://s.yimg.com/os/creatr-uploaded-images/2021-07/7e73d400-df2b-11eb-8b7d-224d93864217.cf.jpg",
		Price:       899,
		Variants: []*Variant{
			{
				Color: "Grey",
				Image: "https://cdn.mos.cms.futurecdn.net/k9Md6R78D8aN4tbGDRWUSE.jpg",
			},
			{
				Color: "Silver",
				Image: "https://i0.wp.com/www.ultimatepocket.com/wp-content/uploads/2021/12/dells-xps-13-9310-is-the-best-13-inch-notebook-you-can-buy-right-now.jpg",
			},
		},
	},
}

var variantsBottle []*Variant = []*Variant{
	{
		Color: "Grey",
		Image: "https://images.unsplash.com/photo-1602143407151-7111542de6e8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8d2F0ZXIlMjBib3R0bGV8ZW58MHx8MHx8&w=1000&q=80",
	},
	{
		Color: "Blue",
		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCzbSXJvxUSmMN9JlQGMy3hVmv_GBkrFl1YQ&usqp=CAU",
	},
}

var variantsLaptop []*Variant = []*Variant{
	{
		Color: "Grey",
		Image: "https://cdn.mos.cms.futurecdn.net/k9Md6R78D8aN4tbGDRWUSE.jpg",
	},
	{
		Color: "Silver",
		Image: "https://i0.wp.com/www.ultimatepocket.com/wp-content/uploads/2021/12/dells-xps-13-9310-is-the-best-13-inch-notebook-you-can-buy-right-now.jpg",
	},
}
