package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"tamaragl/go-birthday-api/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockGetUserUsecaseInterface es un mock para GetUserUsecaseInterface.
type MockGetUserUsecaseInterface struct {
	GetUserBirthdayMessageFunc func(username string) (*entities.BirthdayMessage, error)
}

// GetUserBirthdayMessage implementa la interfaz GetUserUsecaseInterface para el mock.
func (m *MockGetUserUsecaseInterface) GetUserBirthdayMessage(username string) (*entities.BirthdayMessage, error) {
	if m.GetUserBirthdayMessageFunc != nil {
		return m.GetUserBirthdayMessageFunc(username)
	}
	return &entities.BirthdayMessage{}, nil
}

func TestHandleGetHello(t *testing.T) {
	mockUsecase := &MockGetUserUsecaseInterface{}
	server := httptest.NewServer(HandleGetHello(mockUsecase))
	defer server.Close()

	username := "testuser"
	mockMessage := entities.BirthdayMessage{
		Message: "Hello, " + username + "! Your birthday is in 1 day(s)",
	}

	// Modify GetUserBirthdayMessage behavior
	mockUsecase.GetUserBirthdayMessageFunc = func(username string) (*entities.BirthdayMessage, error) {
		return &mockMessage, nil
	}

	// Make request
	req, err := http.NewRequest("GET", server.URL+"/"+username, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check response
	var responseMessage ResponseMessage
	err = json.NewDecoder(resp.Body).Decode(&responseMessage)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, ResponseMessage{Message: mockMessage.Message}, responseMessage)
	assert.NotNil(t, mockUsecase.GetUserBirthdayMessageFunc)
}
