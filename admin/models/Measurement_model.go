package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Measurement struct {
	gorm.Model
	UserID   uint
	Weight   float64
	Height   float64
	FatRatio float64
}

func (measurement Measurement) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&measurement)
}

func (measurement Measurement) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&measurement)
}

func (measurement Measurement) GetAll(where ...interface{}) []Measurement {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var measurements []Measurement
	db.Order("created_at asc").Find(&measurements, where...)
	return measurements
}

func (measurement Measurement) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&measurement, measurement.ID)
}
