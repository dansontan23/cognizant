package main

import (
	"elibrary/config"
	"elibrary/db"
	"elibrary/transport"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.LoadConfig()
	//starting db connection
	db.DbInit()
	defer db.CloseDB()

	http.HandleFunc("/Book", transport.GetBookHandler)
	//http.HandleFunc("/Borrow", transport.BorrowBookHandler)
	//http.HandleFunc("/Extend", transport.ExtendLoanHandler)
	//http.HandleFunc("/Return", transport.ReturnBookHandler)

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
