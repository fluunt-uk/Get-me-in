package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

func (d *Wrapper) GetItem(itemValue string) (*dynamodb.GetItemOutput, error) {

	result, err := d.Connection.GetItem(&dynamodb.GetItemInput{
		TableName: d.Table,
		Key: map[string]*dynamodb.AttributeValue{
			*d.SearchParam: {
				S: aws.String(itemValue),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, &ErrorString{
			Reason: http.StatusText(http.StatusNotFound),
			Code:   http.StatusNotFound,
		}
	}

	return result, nil
}
