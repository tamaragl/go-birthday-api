package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamodbClient struct {
	Client *dynamodb.Client
}

func loadConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}

func NewDynamodbClient(env string) (*DynamodbClient, error) {
	if env == "local" {
		// TODO: Move to a config file by environment
		awsEndpoint := "http://localstack:4566"
		awsRegion := "us-east-1"

		cfgOptions := make([]func(*config.LoadOptions) error, 0)
		// AWS Config
		cfgOptions = append(cfgOptions,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider("dummy", "dummy", "")))
		cfgOptions = append(cfgOptions, config.WithRegion(awsRegion))
		cfg, err := config.LoadDefaultConfig(context.Background(), cfgOptions...)
		if err != nil {
			return nil, fmt.Errorf("initializing AWS session: %s", err)
		}

		clientOpts := make([]func(*dynamodb.Options), 0)

		clientOpts = append(clientOpts, dynamodb.WithEndpointResolver(
			dynamodb.EndpointResolverFromURL(awsEndpoint)))

		return &DynamodbClient{Client: dynamodb.NewFromConfig(cfg, clientOpts...)}, nil
	} else {
		cfg := loadConfig()
		return &DynamodbClient{Client: dynamodb.NewFromConfig(cfg)}, nil
	}
}
