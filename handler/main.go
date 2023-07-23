package main

import (
	"context"
	"koge-api/libs/dynamodb"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := dynamodb.NewDynamoDB("production")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Put(context.TODO(), map[string]string{
		"teste": "aloha",
		"test":  "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello World!",
		StatusCode: 200,
	}, nil
}

func main() {
	if os.Getenv("NODE_ENV") == "local" {
		Handler(context.Background(), events.APIGatewayProxyRequest{
			Body: "", // TODO
		})
	} else {
		lambda.Start(Handler)
	}
}
