package applicant

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golkity/Calc_go/rpn/calc"
)

type Config struct {
	Addr string `json:"addr"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Application struct {
	config *Config
}

func New(configPath string) *Application {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return &Application{config: cfg}
}

func (a *Application) Run() {
	log.Println("CLI mode started. Type 'exit' to quit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("Input expression: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read input:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("Exiting application.")
			return
		}

		result, err := calc.Calc(text)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Printf("Result: %s = %f\n", text, result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		http.Error(w, "Invalid expression: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"expression": req.Expression,
		"result":     result,
	})
}

func (a *Application) RunServer() {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	addr := ":" + a.config.Addr
	log.Printf("Server is running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
