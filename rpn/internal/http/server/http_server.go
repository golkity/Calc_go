package server

import (
	"github.com/golkity/calc_go/rpn/internal/http/handler"
	"log"
	"net/http"
)

func RunHTTPServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.CalculateHandler)

	log.Printf("Сервер запущен тут -> http://localhost%s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
