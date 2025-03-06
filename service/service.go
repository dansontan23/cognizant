package service

import (
	"database/sql"
	"elibrary/dao"
	"elibrary/db"
	"elibrary/models"
)

// creating interface type for easier dependency injection
type Service interface {
	CheckAvailability(title string) (*models.BookDetail, bool)
}

// ServiceImpl is an implementation of Service that includes a database connection
type ServiceImpl struct {
	DB *sql.DB
}

// NewServiceImpl initializes the service with a database connection
func NewServiceImpl() *ServiceImpl {
	// Initialize DB connection
	database := db.DbInit()
	return &ServiceImpl{DB: database}
}

// returns book.model if title is available
func (s *ServiceImpl) CheckAvailability(title string) (*models.BookDetail, bool) {
	return dao.GetBook(s.DB, title)
}
