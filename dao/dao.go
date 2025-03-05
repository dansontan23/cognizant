package dao

import (
	"database/sql"
	"elibrary/db"
	"elibrary/models"
	"log"
)

// basic get function for book
// GetBookHandler handles the request for getting book details
func GetBook(title string) (*models.BookDetail, bool) {

	book, exists := getBookByTitle(title)
	if !exists {
		return nil, false
	}

	return book, true
}

// getBookByTitle queries the SQLite database to retrieve the book by title
func getBookByTitle(title string) (*models.BookDetail, bool) {
	var book models.BookDetail

	// Query the book by title
	row := db.DB.QueryRow("SELECT title, available_copies FROM books WHERE title = ?", title)
	err := row.Scan(&book.Title, &book.AvailableCopies)
	if err != nil {
		if err == sql.ErrNoRows {
			// Book not found
			return nil, false
		}
		log.Println("Error querying book:", err)
		return nil, false
	}

	return &book, true
}

//basic add function for loan

func AddLoan(loan models.LoanDetail) (loans []models.LoanDetail) {
	var Loans = []models.LoanDetail{}
	Loans = append(Loans, loan)
	return Loans

}
