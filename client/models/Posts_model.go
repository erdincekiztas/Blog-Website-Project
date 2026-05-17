package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, PictureUrl string
	CategoryID                                    int
}

func (post Post) Migrate() {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&post)

}

func (post Post) Add() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&post)
}

func (post Post) Get(where ...interface{}) Post {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.First(&post, where...)
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var posts []Post

	db.Find(&posts, where...)

	return posts
}

func (post Post) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&post).Update(column, value)

}

func (post Post) Updates(data Post) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.Model(&post).Updates(data)

}

func (post Post) Delete() {

	//? ilk önce silinecek := admin_models.Post{}.Get() ile bir öğe almalıyız sonrasında diğer fonksiyonlarımızın içerisinde
	//? silinecek.Delete() yaparak öğe silinebilir
	//neden böyle çünkkü gorm ile pointer içerisinde tutuyoruz silinecek değerin struct adresini
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.Delete(&post, post.ID)
}
