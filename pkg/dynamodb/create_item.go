package dynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

//Item created using AttributeValue which is decoded by modeldecoding
func (d *Wrapper) CreateItem(av map[string]*dynamodb.AttributeValue) error {

	// translate into a compatible object
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: d.Table,
	}

	_, errM := d.Connection.PutItem(input)

	if errM != nil {
		log.Println(errM.Error())
		return errM
	}

	return nil
}
