package models

import (
	"github.com/jinzhu/gorm"
	"github.com/meobilivang/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// Instantiate DB instance
	config.Connect()
	db = config.GetDB()

	// migrate Book obj
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// Create new Record in DB
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	// Query for all books
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	// Query with condition & assing to variable `getBook`
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
