package transport

import (
	"database/sql"
	"elibrary/models"
	"elibrary/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBookHandler(t *testing.T) {
	tests := []struct {
		name           string
		title          string
		mockResponse   *models.BookDetail
		err            error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Book not found",
			title:          "Unknown Book",
			mockResponse:   nil,
			err:            sql.ErrNoRows,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `"Book not found"`,
		},
		{
			name:  "Book found",
			title: "Go Programming",
			mockResponse: &models.BookDetail{
				Title:           "Go Programming",
				AvailableCopies: 5,
			},
			err:            nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"Title":"Go Programming","AvailableCopies":5}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.Service)
			mockService.On("CheckAvailability", tt.title).Return(tt.mockResponse, tt.err)

			req, err := http.NewRequest("GET", "/book?title="+tt.title, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Pass mockService here
				GetBookHandler(mockService, w, r)
			})
			handler.ServeHTTP(rr, req)

			// Verify status code
			require.Equal(t, tt.expectedStatus, rr.Code)

			// Verify response body
			require.JSONEq(t, tt.expectedBody, rr.Body.String())

			mockService.AssertExpectations(t)
		})
	}
}
