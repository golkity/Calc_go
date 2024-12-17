package handler

import (
	"encoding/json"
	"errors"
	"github.com/golkity/Calc_go/rpn/Errors"
	"github.com/golkity/Calc_go/rpn/calc"
	"net/http"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, Errors.ErrInvalidExpression) {
			http.Error(w, "Invalid expression", http.StatusUnprocessableEntity)
		} else if errors.Is(err, Errors.ErrDivisionByZero) {
			http.Error(w, "Division by zero", http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CalculateResponse{Result: result})
}
