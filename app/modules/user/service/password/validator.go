package password

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Validator struct {
	MinLength      int
	MaxLength      int
	ForbiddenChars []rune
}

func NewPasswordValidator() *Validator {
	return &Validator{
		MinLength:      8,
		MaxLength:      64,
		ForbiddenChars: []rune{' ', '\t', '\n', '\r', '\'', '"', '\\', '<', '>'},
	}
}

func (pv *Validator) Validate(password string) error {
	length := utf8.RuneCountInString(password)
	if length < pv.MinLength || length > pv.MaxLength {
		return fmt.Errorf("password must be between %d and %d characters", pv.MinLength, pv.MaxLength)
	}

	for _, char := range pv.ForbiddenChars {
		if containsRune(password, char) {
			return fmt.Errorf("forbidden password character `%c`", char)
		}
	}

	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
		return errors.New("password must contain at least one letter")
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	return nil
}

func containsRune(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}
