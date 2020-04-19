package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


func (d *DynamoDB) DeleteItem(itemValue string) error {

	// translate into a compatible object
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			d.SearchParam: {
				S: aws.String(itemValue),
			},
		},
		TableName: aws.String(d.Table),
	}

	_, err := d.Connection.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil
}
