package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// EmailRequest represents the structure of the input email JSON payload
type EmailRequest struct {
	Email string `json:"email"`
}

// Flag represents an individual classification flag for an email
type Flag struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}

// EmailResponse represents the structure of the output JSON response
type EmailResponse struct {
	Email string `json:"email"`
	Flags []Flag `json:"flags"`
}

// Middleware to validate X-RapidAPI-Proxy-Secret header
func validateProxySecret(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxySecret := r.Header.Get("X-RapidAPI-Proxy-Secret")
		expectedSecret := os.Getenv("RAPIDAPI_PROXY_SECRET")

		if proxySecret == "" || proxySecret != expectedSecret {
			http.Error(w, "Invalid or missing X-RapidAPI-Proxy-Secret header", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Validate and classify email
func classifyEmail(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Process email and generate flags
	flags := validateEmail(req.Email)

	// Create and return response
	response := EmailResponse{
		Email: req.Email,
		Flags: flags,
	}
	jsonResponse(w, response, http.StatusOK)
}

// Helper function to validate and classify email
func validateEmail(email string) []Flag {
	flags := []Flag{}

	// Basic email format validation
	emailRegex := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		flags = append(flags, Flag{
			Category:    "syntax",
			Description: "Invalid email format. Example: a@a.com",
		})
		return flags
	}

	// Check for Cyrillic or unusual characters
	for _, r := range email {
		if unicode.Is(unicode.Cyrillic, r) || unicode.IsSymbol(r) || unicode.IsPunct(r) && r != '@' && r != '.' {
			flags = append(flags, Flag{
				Category:    "likely malicious",
				Description: "Email contains Cyrillic or unusual characters.",
			})
			break
		}
	}

	// Check for risky domains
	riskyDomains := []string{".ru", ".su", ".xyz", ".click", ".info", ".top"}
	domain := strings.ToLower(strings.Split(email, "@")[1])
	for _, risky := range riskyDomains {
		if strings.HasSuffix(domain, risky) {
			flags = append(flags, Flag{
				Category:    "spam",
				Description: fmt.Sprintf("Email uses a risky domain: %s", risky),
			})
		}
	}

	// Default to benign if no flags found
	if len(flags) == 0 {
		flags = append(flags, Flag{
			Category:    "benign",
			Description: "No issues detected.",
		})
	}

	return flags
}

// Helper function to send JSON response
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Define routes with middleware
	http.HandleFunc("/classify", validateProxySecret(classifyEmail))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

