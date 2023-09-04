package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"path"
	"tamaragl/go-birthday-api/src/usecases"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func HandleGetHello(usecase usecases.GetUserUsecaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := path.Base(r.URL.String())

		slog.Info("user requested:", "username", username)

		bMsg, err := usecase.GetUserBirthdayMessage(username)
		if err != nil {
			// TODO: validate several types of error
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response := ResponseMessage{Message: bMsg.Message}

		b, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}
