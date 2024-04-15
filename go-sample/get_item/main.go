package main

import (
	"fmt"

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

	input := &dynamodb.GetItemInput{
		TableName: aws.String("test-table"),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String("1"),
			},
		},
	}

	result, err := svc.GetItem(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Item)
}
