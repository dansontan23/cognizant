package dao

import "elibrary/models"

//basic get function for book
func GetBook(title string) (*models.BookDetail, bool) {
	var Books = map[string]*models.BookDetail{}
	book, exists := Books[title]
	return book, exists
}

//basic add function for loan

func AddLoan(loan models.LoanDetail) (loans []models.LoanDetail) {
	var Loans = []models.LoanDetail{}
	Loans = append(Loans, loan)
	return Loans

}
