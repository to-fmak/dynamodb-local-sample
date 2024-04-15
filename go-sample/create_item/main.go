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

	input := &dynamodb.PutItemInput{
		TableName: aws.String("test-table"),
		Item: map[string]*dynamodb.AttributeValue{
			"Id":   {N: aws.String("1")},
			"Name": {S: aws.String("to-fmak")},
		},
	}

	_, err := svc.PutItem(input)
	if err != nil {
		panic(err)
	}
}
