package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"tamaragl/go-birthday-api/src/entities"
	"tamaragl/go-birthday-api/src/usecases"
)

type AddUserInput struct {
	DateOfBirth string `json:"dateOfBirth"`
}

func HandleAddUser(usecase usecases.AddUserUsecaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := path.Base(r.URL.String())
		var input AddUserInput

		// Request
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
			return
		}

		requestUser := entities.User{
			Username:    username,
			DateOfBirth: input.DateOfBirth,
		}

		// Validate
		if v, err := requestUser.IsValid(); err != nil && v == false {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add user
		err = usecase.Add(&requestUser)
		if err != nil {
			http.Error(w, "Error putting item", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
