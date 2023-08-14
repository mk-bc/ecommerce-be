package models

type Variant struct {
	ID        uint `json:"-"`
	Color     string
	Image     string
	ProductID uint `json:"-"`
}
