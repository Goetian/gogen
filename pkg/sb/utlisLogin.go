package sb

import (
	"regexp"
	"strings"
	"unicode"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsValidPassword(passwordToValidate string) (string, bool) {
	//check for leng > 7
	// Upper, Lower, special Char
	var (
		specialChar = "!$#%&^*@€§"
		hasUpper    = false
		hasLower    = false
	)
	if len(passwordToValidate) < 8 {
		return "Password must have a minimum lenght of 8", false
	}
	if strings.Contains(passwordToValidate, specialChar) {
		return "Password must have a minimum lenght of 8", false
	}

	if strings.Contains(passwordToValidate, specialChar) {
		return "Password must have a minimum lenght of 8", false
	}

	for _, charOfPassword := range passwordToValidate {
		if unicode.IsUpper(charOfPassword) {
			hasUpper = true
		}
		if unicode.IsLower(charOfPassword) {
			hasLower = true
		}
	}
	if !hasUpper {
		return "Needs Capital Letter", false
	}

	if !hasLower {
		return "Needs lower Letter", false
	}

	return "Valid Password", true

}
