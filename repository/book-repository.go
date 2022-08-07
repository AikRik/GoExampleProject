package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type BookRepository interface {
	Save(book entity.Book)
	Update(book entity.Book)
	Delete(book entity.Book)
	FindAll() []entity.Book
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewBookRepository() BookRepository {
	db, err := gorm.Open("mysql8", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Book, &entity.Person)

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}

}

func (db *database) Save(book entity.Book) {
	db.connection.Create(&book)
}

func (db *database) Update(book entity.Book) {
	db.connection.Save(&book)
}

func (db *database) Delete(book entity.Book) {
	db.connection.Delete(&book)
}

func (db *database) FindAll() []entity.Book {
	var books []entity.book
	db.connection.Set("gorm:auto_preload", true).Find(&books)

	return books
}
