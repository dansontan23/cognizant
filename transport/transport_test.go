package transport

import (
	"elibrary/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock the service layer
type MockService struct {
	mock.Mock
}

func (m *MockService) CheckAvailability(title string) (*models.BookDetail, bool) {
	args := m.Called(title)
	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}
	return args.Get(0).(*models.BookDetail), args.Bool(1)
}

func TestGetBookHandler(t *testing.T) {
	tests := []struct {
		name           string
		title          string
		mockResponse   *models.BookDetail
		mockExists     bool
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Book not found",
			title:          "Unknown Book",
			mockResponse:   nil,
			mockExists:     false,
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
			mockExists:     true,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"Title":"Go Programming","AvailableCopies":5}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockService)
			mockService.On("CheckAvailability", tt.title).Return(tt.mockResponse, tt.mockExists)

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
