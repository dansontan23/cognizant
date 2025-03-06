package transport

import (
	"elibrary/service"
	"elibrary/utils"
	"log"
	"net/http"
)

// receives request, calls service layer to check for request
func GetBookHandler(s service.Service, w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	book, err := s.CheckAvailability(title)
	if err != nil {
		log.Printf("Error checking availability for book '%s': %v", title, err)
		utils.SendResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	utils.SendResponse(w, http.StatusOK, book)
}
