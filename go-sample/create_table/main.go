package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("ap-northeast-1"),
	}))

	svc := dynamodb.New(sess)

	attrDefinitions := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("Id"),
			AttributeType: aws.String("N"),
		},
	}

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("Id"),
			KeyType:       aws.String("HASH"),
		},
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attrDefinitions,
		TableName:            aws.String("test-table"),
		KeySchema:            keySchema,
		BillingMode:          aws.String("PAY_PER_REQUEST"),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		panic(err)
	}
}
