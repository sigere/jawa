package model

import (
	"github.com/sigere/jawa/db"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string  `query:"name" json:"name" form:"name"`
	Price          float64 `query:"price" json:"price"`
	ManufacturerID uint    `query:"manufacturer_id" json:"manufacturer_id"`
	Manufacturer   Manufacturer
}

func (p *Product) Validate() []string {
	connection := db.GetConnection()
	var result [3]string
	i := 0

	var ids []int
	connection.Table("manufacturers").Where("id = ?", p.ManufacturerID).Pluck("id", &ids)
	if len(ids) != 1 {
		result[i] = "Invalid manufacturer"
		i++
	}

	if p.Name == "" {
		result[i] = "Name cannot be empty"
		i++
	}

	if p.Price <= 0 {
		result[i] = "Price must be > 0"
		i++
	}

	return result[0:i]
}
