package models

import (
	"fmt"
	"reports/services/database"

	"github.com/jinzhu/gorm"
	// For Loading MySQL Drivers
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Book ...
type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// TableName ...
func (b *Book) TableName() string {
	return "books"
}

// GetAllBook ...
func GetAllBook(b *[]Book) (err error) {
	if err = database.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

// AddNewBook ...
func AddNewBook(b *Book) (err error) {
	if err = database.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

// GetOneBook ...
func GetOneBook(b *Book, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

// PutOneBook ...
func PutOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	database.DB.Save(b)
	return nil
}

// DeleteBook ...
func DeleteBook(b *Book, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(b)
	return nil
}
