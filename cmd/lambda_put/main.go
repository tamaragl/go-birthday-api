package main

import (
	"encoding/json"
	"net/http"
	"tamaragl/go-birthday-api/src/entities"
	"tamaragl/go-birthday-api/src/repositories"
	"tamaragl/go-birthday-api/src/storage"
	"tamaragl/go-birthday-api/src/usecases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type AddUserInput struct {
	DateOfBirth string `json:"dateOfBirth"`
}

func PutUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	username := req.PathParameters["username"]

	// Setup
	client, err := storage.NewDynamodbClient("aws")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	repo := repositories.NewDynamodbRepository(client, "Users")
	usecase := usecases.NewAddUserUsecase(repo)

	// Request
	var input AddUserInput
	if err := json.Unmarshal([]byte(req.Body), &input); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: err.Error(),
		}, nil
	}

	requestUser := entities.User{
		Username:    username,
		DateOfBirth: input.DateOfBirth,
	}

	// Validate
	if v, err := requestUser.IsValid(); err != nil && v == false {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	// Add user
	err = usecase.Add(&requestUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(PutUser)
}
