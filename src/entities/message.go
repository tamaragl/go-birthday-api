package entities

import (
	"fmt"
	"time"
)

type BirthdayMessage struct {
	Message string `json:"message"`
}

func NewBirthdayMessage(u *User, t time.Time) (*BirthdayMessage, error) {
	d, err := getDaysUntilBirthday(u.DateOfBirth, t)
	if err != nil {
		return nil, fmt.Errorf("getting days until birthday: %w", err)
	}

	m := getUserMessage(u.Username, d)
	return &BirthdayMessage{Message: m}, nil
}

func getUserMessage(username string, daysUntilBirthday int) string {
	var m string
	if daysUntilBirthday == 0 {
		m = fmt.Sprintf("Hello, %s! It's your birthday today!", username)
	} else {
		m = fmt.Sprintf("Hello, %s! Your birthday is in %d day(s)", username, daysUntilBirthday)
	}

	return m
}

func getDaysUntilBirthday(dateOfBirth string, now time.Time) (int, error) {
	dob, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		return -1, fmt.Errorf("parsing dateOfBirth: %w", err)
	}

	currentYear := now.Year()
	nextBirthday := time.Date(now.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)

	// If the birthday was in the past of this yeas, it add 1 year
	if now.After(nextBirthday) {
		nextBirthday = time.Date(currentYear+1, dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)
	}

	daysUntilBirthday := int(nextBirthday.Sub(now).Hours() / 24) // convert to days

	return daysUntilBirthday, nil
}
