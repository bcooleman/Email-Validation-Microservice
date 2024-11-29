package emailvalidator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

// EmailRequest represents the incoming JSON payload
type EmailRequest struct {
	Email string `json:"email"`
}

// Flag represents a validation issue with the email
type Flag struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}

// EmailResponse represents the response structure
type EmailResponse struct {
	Email string `json:"email"`
	Flags []Flag `json:"flags"`
}

// ValidateEmail is the HTTP Cloud Function entry point
func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	flags := performValidation(req.Email)
	response := EmailResponse{
		Email: req.Email,
		Flags: flags,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func performValidation(email string) []Flag {
	flags := []Flag{}

	// Basic email format validation
	emailRegex := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		flags = append(flags, Flag{
			Category:    "syntax",
			Description: "Invalid email format. Example: a@b.com",
		})
		return flags
	}

	// Check for Cyrillic or unusual characters
	for _, r := range email {
		if unicode.Is(unicode.Cyrillic, r) || unicode.IsSymbol(r) && r != '@' && r != '.' {
			flags = append(flags, Flag{
				Category:    "malicious",
				Description: "Contains Cyrillic or unusual characters.",
			})
			break
		}
	}

	// Risky domains
	riskyDomains := []string{".ru", ".xyz", ".click", ".info", ".su"}
	domain := strings.Split(strings.ToLower(email), "@")[1]
	for _, risky := range riskyDomains {
		if strings.HasSuffix(domain, risky) {
			flags = append(flags, Flag{
				Category:    "spam",
				Description: fmt.Sprintf("Email uses a risky domain: %s", risky),
			})
		}
	}

	if len(flags) == 0 {
		flags = append(flags, Flag{
			Category:    "benign",
			Description: "No issues detected.",
		})
	}

	return flags
}

