package main

import (
	"elibrary/config"
	"elibrary/transport"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/Book", transport.GetBookHandler)
	//http.HandleFunc("/Borrow", transport.BorrowBookHandler)
	//http.HandleFunc("/Extend", transport.ExtendLoanHandler)
	//http.HandleFunc("/Return", transport.ReturnBookHandler)

	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
