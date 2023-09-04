package handler

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"tamaragl/go-birthday-api/src/entities"
	"testing"
)

// MockAddUserUsecase
type MockAddUserUsecase struct {
	AddCalled bool
}

func (m *MockAddUserUsecase) Add(user *entities.User) error {
	m.AddCalled = true
	return nil
}

func TestHandleAddUser(t *testing.T) {
	mockUsecase := &MockAddUserUsecase{}

	server := httptest.NewServer(HandleAddUser(mockUsecase))
	defer server.Close()

	// Test valid user
	validUserData := `{"dateOfBirth": "2000-02-15"}`
	slog.Info("server:", "url:", server.URL)
	req, err := http.NewRequest(http.MethodPut, server.URL+"/Pepita", bytes.NewBufferString(validUserData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
	}

	if !mockUsecase.AddCalled {
		t.Error("Add method not called on usecase")
	}

	// Test invalid username
	invalidUserData := `{"dateOfBirth": "2000-02-15"}`
	req, err = http.NewRequest(http.MethodPut, server.URL+"/123", bytes.NewBufferString(invalidUserData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	// Test invalid date of birth
	invalidBirth := `{"username": "abcd", "dateOfBirth": "3000-02-15"}`
	req, err = http.NewRequest(http.MethodPut, server.URL, bytes.NewBufferString(invalidBirth))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
