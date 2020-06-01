package dynamodb

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io"
)

/**
* Convert type interface to dynamodb readable object and JSON
**/
func DecodeToDynamoAttribute(readBody io.ReadCloser, m interface{}) (map[string]*dynamodb.AttributeValue, error) {

	if err := DecodeToMap(readBody, &m); err != nil {
		return nil, err
	}

	av, errM := dynamodbattribute.MarshalMap(&m)

	if errM != nil {
		return nil, errM
	}

	return av, nil
}

/**
* Convert the interface fields into a map
**/
func DecodeToMap(b io.ReadCloser, m interface{})  error {

	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errJson := json.NewDecoder(b).Decode(&m)

	if errJson != nil {
		return errJson
	}

	return nil
}

/**
* Model mapping of type interface to item from dynamodb
**/
func Unmarshal(result *dynamodb.GetItemOutput, m interface{}) error {

	err := dynamodbattribute.UnmarshalMap(result.Item, &m)

	if err != nil {
		return err
	}

	return nil
}

func AddEmptyCollection(av map[string]*dynamodb.AttributeValue, v string){

	av[v].NULL = nil
	av[v].M = map[string]*dynamodb.AttributeValue{}
}