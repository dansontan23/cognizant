package dao

import (
	"database/sql"
	"elibrary/models"
	"log"
)

// basic get function for book
// GetBookHandler handles the request for getting book details
func GetBook(db *sql.DB, title string) (*models.BookDetail, bool) {

	book, exists := getBookByTitle(db, title)
	if !exists {
		return nil, false
	}

	return book, true
}

//basic add function for loan

func AddLoan(loan models.LoanDetail) (loans []models.LoanDetail) {
	var Loans = []models.LoanDetail{}
	Loans = append(Loans, loan)
	return Loans

}

// helper function, small caps to not expose func
// getBookByTitle queries the SQLite database to retrieve the book by title
func getBookByTitle(db *sql.DB, title string) (*models.BookDetail, bool) {
	var book models.BookDetail
	log.Println("Quering db")
	// Query the book by title
	row := db.QueryRow("SELECT title, available_copies FROM books WHERE title = ?", title)
	err := row.Scan(&book.Title, &book.AvailableCopies)
	if err != nil {
		log.Println("Error querying book:", err)
		return nil, false
	}
	log.Println("Successfully queried book")
	return &book, true
}
