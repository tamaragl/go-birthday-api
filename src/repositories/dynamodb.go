package repositories

import (
	"context"
	"fmt"
	"log/slog"
	"tamaragl/go-birthday-api/src/entities"
	"tamaragl/go-birthday-api/src/storage"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamodbRepository struct {
	c     *storage.DynamodbClient
	table string
}

func NewDynamodbRepository(client *storage.DynamodbClient, table string) *DynamodbRepository {
	return &DynamodbRepository{c: client, table: table}
}

func (r *DynamodbRepository) GetItem(item string) (*entities.User, error) {
	slog.Info("getItem:", "username", item)

	user := &entities.User{}

	// the composite primary key of the movie in a format that can be
	// sent to DynamoDB.
	username, err := attributevalue.Marshal(item)
	if err != nil {
		panic(err)
	}
	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{"Username": username}, TableName: aws.String(r.table),
	}

	// GetItem
	out, err := r.c.Client.GetItem(context.TODO(), getItemInput)
	if err != nil {
		slog.Info("dynamodb error:", "error", err)
		return user, fmt.Errorf("dynamoDB error: %s", err)
	}

	slog.Info("output:", "output", out)

	// TODO: Move unmarshall to usercase
	// Unmarshall getItemOutput to User
	if len(out.Item) > 0 {
		if err := attributevalue.UnmarshalMap(out.Item, user); err != nil {
			return user, fmt.Errorf("unmarshalling AttributeValue to user: %w", err)
		}
	} else {
		return user, fmt.Errorf("user not found")
	}

	return user, nil
}

func (r *DynamodbRepository) PutItem(user *entities.User) error {
	item, err := attributevalue.MarshalMap(user)
	_, err = r.c.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(r.table), Item: item,
	})
	if err != nil {
		return fmt.Errorf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return nil
}
