package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/matcornic/hermes/v2" // SMTP check (optional, configure as needed)
)

type ValidationResult struct {
	Email            string `json:"email"`
	SyntaxValid      bool   `json:"syntax_valid"`
	DomainValid      bool   `json:"domain_valid"`
	DisposableDomain bool   `json:"disposable_domain"`
	RiskyDomain      bool   `json:"risky_domain"`
	OnBlacklist      bool   `json:"on_blacklist"`
	ContainsCyrillic bool   `json:"contains_cyrillic"`
	SMTPValid        bool   `json:"smtp_valid"`
}

var (
	disposableDomains = []string{"mailinator.com", "trashmail.com", "tempmail.com"} // Add more as needed
	riskyDomains      = []string{".ru", ".xyz", ".click", ".info", ".su"}
	blacklistDomains  = []string{"spamdomain.com", "blacklisteddomain.xyz"}
)

func validateEmailHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request map[string]string
	if err := decoder.Decode(&request); err != nil || request["email"] == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := request["email"]
	result := validateEmail(email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func validateEmail(email string) ValidationResult {
	return ValidationResult{
		Email:            email,
		SyntaxValid:      validateSyntax(email),
		DomainValid:      validateDomain(email),
		DisposableDomain: checkDisposableDomain(email),
		RiskyDomain:      checkRiskyDomain(email),
		OnBlacklist:      checkBlacklist(email),
		ContainsCyrillic: containsCyrillic(email),
		SMTPValid:        validateSMTP(email),
	}
}

func validateSyntax(email string) bool {
	emailRegex := `^[^\s@]+@[^\s@]+\.[^\s@]+$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func validateDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	_, err := net.LookupMX(domain)
	return err == nil
}

func checkDisposableDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	for _, disposable := range disposableDomains {
		if domain == disposable {
			return true
		}
	}
	return false
}

func checkRiskyDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	for _, risky := range riskyDomains {
		if strings.HasSuffix(domain, risky) {
			return true
		}
	}
	return false
}

func checkBlacklist(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	for _, blacklist := range blacklistDomains {
		if domain == blacklist {
			return true
		}
	}
	return false
}

func containsCyrillic(email string) bool {
	for _, r := range email {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

func validateSMTP(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}

	// Simple SMTP check (Optional, as some servers may block this)
	mx := mxRecords[0].Host
	conn, err := net.Dial("tcp", mx+":25")
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	http.HandleFunc("/validate", validateEmailHandler)

	port := "8080"
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
 
