package models

type BookDetail struct {
	Title           string
	AvailableCopies int
}

var Books = map[string]*BookDetail{}
