package main

import (
	"fmt"
	"github.com/golkity/calc_go/config"
	"github.com/golkity/calc_go/rpn/internal/http/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	port := cfg.Port
	addr := fmt.Sprintf(":%s", port)
	server.RunHTTPServer(addr)
}
