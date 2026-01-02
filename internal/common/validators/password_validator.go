package validators

import (

	"unicode"
	"strings"

) 


func containsIgnoreCase(s, substr string) bool {
    return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}


func PasswordStrengthScore(password string) int {
    score := 0
    length := len(password)

    // Length score
    if length >= 8 {
        score += 10
    }
    if length >= 12 {
        score += 10
    }
    if length >= 16 {
        score += 10
    }

    hasUpper := false
    hasLower := false
    hasDigit := false
    hasSpecial := false

    for _, ch := range password {
        switch {
        case unicode.IsUpper(ch):
            hasUpper = true
        case unicode.IsLower(ch):
            hasLower = true
        case unicode.IsDigit(ch):
            hasDigit = true
        case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
            hasSpecial = true
        }
    }

    if hasUpper {
        score += 15
    }
    if hasLower {
        score += 15
    }
    if hasDigit {
        score += 15
    }
    if hasSpecial {
        score += 15
    }

    // Variety bonus
    varietyCount := 0
    if hasUpper {
        varietyCount++
    }
    if hasLower {
        varietyCount++
    }
    if hasDigit {
        varietyCount++
    }
    if hasSpecial {
        varietyCount++
    }

    if varietyCount >= 3 {
        score += 10
    }

    // Penalty for common weak patterns
    weakPatterns := []string{
        "password", "123456", "qwerty", "admin", "welcome",
    }

    for _, p := range weakPatterns {
        if containsIgnoreCase(password, p) {
            score -= 20
            break
        }
    }

    if score < 0 {
        score = 0
    }
    if score > 100 {
        score = 100
    }

    return score
}
