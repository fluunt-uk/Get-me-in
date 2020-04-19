package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Item created using AttributeValue which is decoded by modeldecoding
func (d *DynamoDB) CreateItem(av map[string]*dynamodb.AttributeValue) error {

	// translate into a compatible object
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.Table),
	}

	_, errM := d.Connection.PutItem(input)

	if errM != nil {
		return errM
	}

	return nil
}
