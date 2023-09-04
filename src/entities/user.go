package entities

import (
	"fmt"
	"regexp"
	"time"
)

type User struct {
	Username    string `json:"username"    dynamodbav:"Username"`
	DateOfBirth string `json:"dateOfBirth"    dynamodbav:"DateOfBirth"`
}

func (u *User) IsValid() (bool, error) {
	if !validateUsername(u.Username) {
		return false, fmt.Errorf("the username '%s' is invalid, must contain only letters", u.Username)
	}
	if !validateDateOfBirth(u.DateOfBirth) {
		return false, fmt.Errorf("the date of birth  '%s' is invalid, must be this format YYYY-MM-DD and be a date before the today date.", u.DateOfBirth)
	}

	return true, nil
}

func validateUsername(username string) bool {
	match, _ := regexp.MatchString("^[A-Za-z]+$", username)
	return match
}

func validateDateOfBirth(dateOfBirth string) bool {
	dob, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		return false
	}

	now := time.Now()

	return dob.Before(now)
}
