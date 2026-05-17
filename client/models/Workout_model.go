package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserID     uint
	DayID      int
	Exercise   string
	Sets       string
	Reps       string
	RestPeriod string

	// antrenman detayları, egzersizler vs.
}

func (workout Workout) Migrate() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&workout)
}

func (workout Workout) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&workout)
}

func (workout Workout) Get(where ...interface{}) Workout {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.First(&workout, where...)
	return workout
}

func (workout Workout) GetAll(where ...interface{}) []Workout {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var workouts []Workout
	db.Find(&workouts, where...)
	return workouts
}

func (workout Workout) Updates(data Workout) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&workout).Updates(data)
}

func (workout Workout) Delete() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&workout, workout.ID)
}
