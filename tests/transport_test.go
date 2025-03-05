package transport

import (
	"elibrary/transport"
	"net/http"
	"net/http/httptest"
	"testing"
)

// testing get function based on error code
func TestGetBookHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/Book?title=Sample", nil)
	resp := httptest.NewRecorder()
	transport.GetBookHandler(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", resp.Code)
	}
}
