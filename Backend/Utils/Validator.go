package utils

import (
	"errors"
	"regexp"
	"strings"
)

type Validator struct{}

type ValidatorInterface interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
	ValidateRole(role string) error
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func (v *Validator) ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password must be between 8 and 30 characters long")
	}

	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString
	)

	if !hasUpper(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber(password) {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (v *Validator) ValidateRole(role string) error {
	if strings.ToLower(role) != "educator" && strings.ToLower(role) != "student"{
		return errors.New("role must be either educator or student")
	}
	return nil
}