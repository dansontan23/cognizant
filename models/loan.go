package models

type LoanDetail struct {
	NameOfBorrower string
	LoanDate       string
	ReturnDate     string
}

var Loans = []LoanDetail{}
