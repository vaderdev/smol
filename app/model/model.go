package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Smol struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json: "redirect" gorm:"not null"`
	Smol     string `json: "smol", gorm: "unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	dsn := "host=localhost user=postgres password=postgres dbname=smol port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Smol{})
	if err != nil {
		fmt.Println(err)
	}
}
