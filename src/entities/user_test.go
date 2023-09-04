package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTrue(t *testing.T) {
	u := User{
		Username:    "test",
		DateOfBirth: "1990-01-01",
	}

	v, _ := u.IsValid()

	assert.True(t, v)
}

func TestIsValidUsernameFalse(t *testing.T) {
	u := User{
		Username:    "123",
		DateOfBirth: "1990-01-01",
	}

	v, _ := u.IsValid()

	assert.False(t, v)
}

func TestIsValidDateOfBirthFalse(t *testing.T) {
	u := User{
		Username:    "test",
		DateOfBirth: "30-01-01",
	}

	v, _ := u.IsValid()

	assert.False(t, v)
}

func TestValidateUsername(t *testing.T) {
	match := validateUsername("MariCarmen")
	assert.True(t, match)
}

func TestValidateDateOfBirthToday(t *testing.T) {
	dob := time.Now().Format("2006-01-02")
	match := validateDateOfBirth(dob)
	assert.True(t, match)
}

func TestValidateDateOfBirthInvalid(t *testing.T) {
	dob := time.Now().Format("3006-19-02")
	match := validateDateOfBirth(dob)
	assert.False(t, match)
}

func TestValidateDateOfBirthErrorParser(t *testing.T) {
	dob := time.Now().Format("3006-19-SS")
	match := validateDateOfBirth(dob)
	assert.False(t, match)
}
