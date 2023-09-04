#!/bin/bash

echo "########### Executing create_dynamodb.sh ###########"

echo "---------- Table creation ----------"
awslocal dynamodb create-table \
    --table-name Users \
    --attribute-definitions AttributeName=Username,AttributeType=S \
    --key-schema AttributeName=Username,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

echo "---------- Table data ----------"
# echo $PWD
cd /etc/localstack/init/ready.d/
awslocal dynamodb batch-write-item --request-items file://items.json
