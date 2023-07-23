package main

import (
	"context"
	"fmt"
	"koge-api/libs/dynamodb"
)

func main() {
	fmt.Println("hi, i'm working! :)")

	db, err := dynamodb.NewDynamoDB("local")
	if err != nil {
		panic(err)
	}

	db.Put(context.TODO(), map[string]string{
		"teste": "aloha",
		"test":  "123",
	})
}
