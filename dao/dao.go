package dao

import (
	"database/sql"
	"elibrary/models"
	"log"
)

// creating interface type for easier dependency injection
type Dao interface {
	GetBook(db *sql.DB, title string) (*models.BookDetail, error)
	// AddLoan(db *sql.DB, loan models.LoanDetail) (loans []models.LoanDetail)
}

// DaoImpl implements the Dao interface
type DaoImpl struct{}

// NewDao creates a new instance of DaoImpl
func NewDao() Dao {
	return &DaoImpl{}
}

// basic get function for book
// GetBookHandler handles the request for getting book details
func (d *DaoImpl) GetBook(db *sql.DB, title string) (*models.BookDetail, error) {

	book, err := getBookByTitle(db, title)
	if err != nil {
		return nil, err
	}

	return book, err
}

//basic add function for loan

// func (d *DaoImpl) AddLoan(db *sql.DB, loan models.LoanDetail) []models.LoanDetail {
// 	var Loans = []models.LoanDetail{}
// 	Loans = append(Loans, loan)
// 	return Loans

// }

// helper function, small caps to not expose func
// getBookByTitle queries the SQLite database to retrieve the book by title
func getBookByTitle(db *sql.DB, title string) (*models.BookDetail, error) {
	var book models.BookDetail
	log.Println("Quering db")
	// Query the book by title
	row := db.QueryRow("SELECT title, available_copies FROM books WHERE title = ?", title)
	err := row.Scan(&book.Title, &book.AvailableCopies)

	if err != nil {
		if err == sql.ErrNoRows {
			// Log the "no rows" scenario and return a book with default values
			log.Println("No rows found for book:", title, err)
			return nil, err
		}
		log.Println("Error querying book:", err)
		return nil, err
	}
	log.Println("Successfully queried book")
	return &book, nil
}
