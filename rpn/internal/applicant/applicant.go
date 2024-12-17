package applicant

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golkity/calc_go/internal/calc"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}
	return &Config{
		Addr: addr,
	}
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) Run() error {
	log.Println("CLI mode started. Type 'exit' to quit.")
	for {
		log.Print("Input expression: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read expression from console:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("Application successfully closed.")
			return nil
		}

		result, err := calc.Calc(text)
		if err != nil {
			log.Printf("Failed to calculate expression '%s': %v", text, err)
		} else {
			log.Printf("Result: %s = %f", text, result)
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

	var request Request
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, errors.New("next")) {
			http.Error(w, "Invalid expression: "+err.Error(), http.StatusUnprocessableEntity)
		} else {
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"expression": request.Expression,
		"result":     result,
	})
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	addr := ":" + a.config.Addr
	log.Printf("Server is running on http://localhost%s", addr)
	return http.ListenAndServe(addr, nil)
}
