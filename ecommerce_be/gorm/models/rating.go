package models

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model `json:"-"`
	Name       string
	Review     string
	Rating     uint
	ProductID  uint `json:"-"`
}
