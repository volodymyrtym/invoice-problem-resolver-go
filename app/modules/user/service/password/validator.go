package password

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Validator struct {
	minLength      int
	maxLength      int
	forbiddenChars []rune
}

func NewPasswordValidator() *Validator {
	return &Validator{
		minLength:      8,
		maxLength:      64,
		forbiddenChars: []rune{' ', '\t', '\n', '\r', '\'', '"', '\\', '<', '>'},
	}
}

func (pv *Validator) Validate(password string) error {
	length := utf8.RuneCountInString(password)
	if length < pv.minLength || length > pv.maxLength {
		return fmt.Errorf("password must be between %d and %d characters", pv.minLength, pv.maxLength)
	}

	for _, char := range pv.forbiddenChars {
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
