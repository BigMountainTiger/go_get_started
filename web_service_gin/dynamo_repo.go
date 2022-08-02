package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoItem struct {
	Id         string `json:"id"`
	Entry_time string `json:"entry_time"`
	Attr       string `json:"attr"`
}

var client = dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
})))

func get_dynamo_entry() (*DynamoItem, error) {

	tableName := "TABLE_1"

	id := "2"
	entry_time := "entry_2"

	result, err := client.GetItem(&dynamodb.GetItemInput{
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
		return nil, err
	}

	item := DynamoItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
