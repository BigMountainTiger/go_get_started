package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	id         string
	entry_time string
	attr       string
}

var svc = dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
})))

func get_item(id string, entry_time string) {
	var tableName = "TABLE_1"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"entry_time": {
				S: aws.String(entry_time),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println(result.Item)
	fmt.Println(item)

	var ditem = result.Item

	fmt.Println(ditem["attr"])
}

func main() {
	get_item("2", "entry_2")
}
