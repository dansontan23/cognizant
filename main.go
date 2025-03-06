package main

import (
	"elibrary/config"
	"elibrary/dao"
	"elibrary/handler"
	"elibrary/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.LoadConfig()

	//starting db connection by calling service layer
	daoImpl := dao.NewDao()
	services := service.NewServiceImpl(daoImpl)
	services.DB.SetConnMaxIdleTime(5)
	//closing connection if ending abruptly
	defer services.DB.Close()

	// Initialize Mux router
	router := mux.NewRouter()
	// Setup HTTP routes
	handler.SetupHandlers(services, router)
	log.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
