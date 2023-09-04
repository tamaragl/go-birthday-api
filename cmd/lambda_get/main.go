package main

import (
	"encoding/json"
	"net/http"
	"tamaragl/go-birthday-api/src/repositories"
	"tamaragl/go-birthday-api/src/storage"
	"tamaragl/go-birthday-api/src/usecases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getHello(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	usecase := usecases.NewGetUserUsecase(repo)

	// Get birthday message
	bMsg, err := usecase.GetUserBirthdayMessage(username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       err.Error(),
		}, nil
	}

	jsonData, _ := json.Marshal(bMsg)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonData),
	}, nil

}

func main() {
	lambda.Start(getHello)
}
