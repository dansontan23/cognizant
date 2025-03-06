package transport

import (
	"elibrary/service"
	"elibrary/utils"
	"net/http"
)

// receives request, calls service layer to check for request
func GetBookHandler(s service.Service, w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	book, exists := s.CheckAvailability(title)
	if !exists {
		utils.SendResponse(w, http.StatusNotFound, "Book not found")
		return
	}
	utils.SendResponse(w, http.StatusOK, book)
}
