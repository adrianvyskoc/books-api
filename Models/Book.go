package Models

import (
	"books/Config"
)

func GetAllBooks(book *[]Book) (err error) {
	if err = Config.DB.Find(book).Error; err != nil {
		return err
	}
	return nil
}

func CreateBook(book *Book) (err error) {
	if err = Config.DB.Create(book).Error; err != nil {
	 return err
	}
	return nil
}

func DeleteBookByID(id string) (err error) {
	if err = Config.DB.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}