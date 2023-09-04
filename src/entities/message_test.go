package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayMessage(t *testing.T) {
	msg, _ := NewBirthdayMessage(&User{Username: "TestUsername", DateOfBirth: "2016-02-01"}, time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, "Hello, TestUsername! It's your birthday today!", msg.Message)
}

func TestNewBirthdayMessageErrorDays(t *testing.T) {
	_, err := NewBirthdayMessage(&User{Username: "TestUsername", DateOfBirth: "2016-02-130"}, time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestGetUserMessageTodayBirthday(t *testing.T) {
	message := getUserMessage("TestUsername", 0)
	expectedMsg := "Hello, TestUsername! It's your birthday today!"

	assert.Equal(t, expectedMsg, message)
}

func TestGetUserMessageFutureBirthday(t *testing.T) {
	message := getUserMessage("TestUsername", 3)
	expectedMsg := "Hello, TestUsername! Your birthday is in 3 day(s)"

	assert.Equal(t, expectedMsg, message)
}

func TestGetDaysUntilBirthdayBeenToday(t *testing.T) {
	now := time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)
	daysUntilBirthday, _ := getDaysUntilBirthday("2016-02-01", now)

	assert.Equal(t, 0, daysUntilBirthday)
}

func TestGetDaysUntilBirthdayBeenTomorrow(t *testing.T) {
	now := time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)
	daysUntilBirthday, _ := getDaysUntilBirthday("2016-02-02", now)

	assert.Equal(t, 1, daysUntilBirthday)
}

func TestGetDaysUntilBirthdayFuture(t *testing.T) {
	dateOfBirth := "2016-02-15"
	now := time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)
	daysUntilBirthday, _ := getDaysUntilBirthday(dateOfBirth, now)

	assert.Equal(t, 14, daysUntilBirthday)
}

func TestGetDaysUntilBirthdayPast(t *testing.T) {
	dateOfBirth := "2016-02-15"
	now := time.Date(2016, 2, 20, 0, 0, 0, 0, time.UTC)
	daysUntilBirthday, _ := getDaysUntilBirthday(dateOfBirth, now)

	assert.Equal(t, 361, daysUntilBirthday)
}

func TestGetDaysUntilBirthdayError(t *testing.T) {
	dateOfBirth := "2016-02-150"
	now := time.Now()
	_, err := getDaysUntilBirthday(dateOfBirth, now)

	assert.Error(t, err)
}
