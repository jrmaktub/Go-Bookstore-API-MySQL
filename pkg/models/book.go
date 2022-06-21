package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jrmaktub/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	//calling the function on the config file to initialize DB
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//b is parameter, and we return is  also book
func (b *Book) CreateBook() *Book {
	//gorm handles queries as an ORM
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	//running an SQL command
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

//controll starts with routes
//routes give  control to  controllers
//controller gives controll to Book

//user interacts with routes
//and routes send controll to controller logic
//controller makes operation with DB
