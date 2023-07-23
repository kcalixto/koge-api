package main

import (
	"context"
	"koge-api/libs/dynamodb"
)

func main() {
	db, err := dynamodb.NewDynamoDB("local")
	if err != nil {
		panic(err)
	}

	db.Put(context.TODO(), map[string]string{
		"teste": "aloha",
		"test":  "123",
	})
}
