package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct untuk  menerima input
type Request struct {
	Operation string  `json:"operation"`
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
}

// struct untuk mengirim hasil
type Response struct {
	Result float64 `json:"result"`
}

// middleware untuk mengatur CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Access-Control-Allow-Credentials")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func calculate(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res Response

	// decode json dari request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// calculate based on operation
	switch req.Operation {
	case "add":
		res.Result = req.Operand1 + req.Operand2
	case "subtract":
		res.Result = req.Operand1 - req.Operand2
	case "multiply":
		res.Result = req.Operand1 * req.Operand2
	case "divide":
		if req.Operand2 == 0 {
			http.Error(w, "Cant divide by zero", http.StatusBadRequest)
			return
		}
		res.Result = req.Operand1 / req.Operand2
	default:
		http.Error(w, "Unknown operation", http.StatusBadRequest)
		return
	}

	// encode ke json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.Handle("/calculate", enableCORS(http.HandlerFunc(calculate)))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
