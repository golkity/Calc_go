package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestingCaclHandler(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		status  int
		expBody string
	}{
		{
			name:    "Valid Expression",
			input:   `{"expression": "2+2*2"}`,
			status:  http.StatusOK,
			expBody: `{"result":6}`,
		},
		{
			name:    "Division by Zero",
			input:   `{"expression": "4/0"}`,
			status:  http.StatusUnprocessableEntity,
			expBody: "Division by zero",
		},
		{
			name:    "Invalid Expression",
			input:   `{"expression": "2+*3"}`,
			status:  http.StatusUnprocessableEntity,
			expBody: "Invalid expression",
		},
		{
			name:    "Invalid JSON",
			input:   `{"expr": "2+2"}`,
			status:  http.StatusBadRequest,
			expBody: "Invalid JSON format",
		},
		{
			name:    "Invalid HTTP Method",
			input:   "",
			status:  http.StatusMethodNotAllowed,
			expBody: "Only POST method is allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.input))
			if tt.name == "Invalid HTTP Method" {
				req = httptest.NewRequest(http.MethodGet, "/", nil)
			}

			rr := httptest.NewRecorder()
			CalculateHandler(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("Expected status %d, got %d", tt.status, status)
			}

			if body := rr.Body.String(); !bytes.Contains([]byte(body), []byte(tt.expBody)) {
				t.Errorf("Expected body to contain %q, got %q", tt.expBody, body)
			}
		})
	}
}
