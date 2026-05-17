package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Diet struct {
	gorm.Model
	UserID   uint
	MealID   int
	Food     string
	Grammage string
}

func (diet Diet) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&diet)
}

func (diet Diet) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&diet)
}

func (diet Diet) Get(where ...interface{}) Diet {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.First(&diet, where...)
	return diet
}

func (diet Diet) GetAll(where ...interface{}) []Diet {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var diets []Diet
	db.Find(&diets, where...)
	return diets
}

func (diet Diet) Updates(data Diet) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&diet).Updates(data)
}

func (diet Diet) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&diet, diet.ID)
}
