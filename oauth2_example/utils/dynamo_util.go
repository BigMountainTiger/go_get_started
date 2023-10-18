package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type History struct {
	Date   string
	Action string
}

type DynamoItem struct {
	Loyalty_eligibility        bool
	Loyalty_enrollment_status  bool
	Loyalty_tier               string
	Customer_id                string
	Loyalty_enrollment_history []History
}

func Get_loyalty_profile(customer_id string) (*DynamoItem, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "saml",
	}))
	sess.Config.Region = aws.String("us-east-1")
	var client = dynamodb.New(sess)

	tableName := "loyalty-profile-global-np"
	result, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customer_id),
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
