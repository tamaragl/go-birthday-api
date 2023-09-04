package main

import (
	"log"
	"net/http"
	"tamaragl/go-birthday-api/cmd/server_http/handler"
	"tamaragl/go-birthday-api/src/repositories"
	"tamaragl/go-birthday-api/src/storage"
	"tamaragl/go-birthday-api/src/usecases"

	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	client, err := storage.NewDynamodbClient("local")
	if err != nil {
		log.Fatalf("dynamodb client: %s", err)
	}
	repo := repositories.NewDynamodbRepository(client, "Users")
	getUser := usecases.NewGetUserUsecase(repo)
	addUser := usecases.NewAddUserUsecase(repo)

	r := mux.NewRouter()
	r.HandleFunc("/hello/{username}", handler.HandleGetHello(getUser)).Methods(http.MethodGet)
	r.HandleFunc("/hello/{username}", handler.HandleAddUser(addUser)).Methods(http.MethodPut)

	return r
}
