package model

import (
	"github.com/sigere/jawa/db"
	"gorm.io/gorm"
	"time"
)

func Migrate(connection *db.Connection) error {
	return connection.AutoMigrate(
		&Product{},
		&Manufacturer{},
		&Order{},
		&Bundle{},
	)
}

type Manufacturer struct {
	gorm.Model
	Name          string
	Address       string
	LoyaltyPoints uint16
}

type Order struct {
	gorm.Model
	CompletedAt time.Time
	Bundles     []Bundle
}

type Bundle struct {
	gorm.Model
	Product   Product
	ProductID uint
	Amount    uint
	OrderID   uint
}
