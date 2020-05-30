package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"net/http"
	"os"
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

func (d *Wrapper) GetAll(queryBy string) ([]interface{}, error) {

	// Create the Expression to fill the input struct with.
	filt := expression.Name("field").Equal(expression.Value(queryBy))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 d.Table,
	}

	// Make the DynamoDB Query API call
	result, err := d.Connection.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, &ErrorString{
			Reason: http.StatusText(http.StatusNotFound),
			Code:   http.StatusNotFound,
		}
	}

	var jsonArray []interface{}
	for _,i := range result.Items {

		m := *d.GenericModel
		err = dynamodbattribute.UnmarshalMap(i, &m)
		jsonArray = append(jsonArray, m)
	}

	return jsonArray, nil
}