package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

const MaxInputLength = 30

var inputRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func main() {
	http.HandleFunc("/input", handleInput)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("user_input")

	if !validateInput(input) {
		http.Error(w, "La entrada es demasiado larga. Por favor, ingrese una entrada de menos de 30 caracteres.", http.StatusBadRequest)
		return
	}

	sanitizedInput := sanitizeInput(input)

	fmt.Fprintf(w, "Processed input: %s", sanitizedInput)
}

func validateInput(input string) bool {
	return utf8.RuneCountInString(input) <= MaxInputLength && inputRegex.MatchString(input)
}

func sanitizeInput(input string) string {
	return strings.TrimSpace(input)
}