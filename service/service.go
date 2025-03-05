package service

import (
	"elibrary/dao"
	"elibrary/models"
)

// returns book.model if title is available
func CheckAvailability(title string) (*models.BookDetail, bool) {
	return dao.GetBook(title)
}
