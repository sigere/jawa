package main

import "gorm.io/gorm"

func migrate(connection *Connection) error {
	return connection.AutoMigrate(
		&Product{},
		&Manufacturer{},
	)
}

type Product struct {
	gorm.Model
	Name           string
	Price          float64
	ManufacturerID uint
	Manufacturer   Manufacturer
}

type Manufacturer struct {
	gorm.Model
	Name          string
	Address       string
	LoyaltyPoints uint16
}
