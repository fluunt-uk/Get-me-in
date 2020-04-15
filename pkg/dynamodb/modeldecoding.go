package dynamodb

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io"
)

/**
* Convert type interface to dynamodb readable object and JSON
**/
func DecodeToDynamoAttribute(readBody io.ReadCloser, m interface{}) (map[string]*dynamodb.AttributeValue, error) {

	bodyMap, err := DecodeToMap(readBody, &m)

	if err != nil {
		return nil, err
	}

	//encoder := dynamodbattribute.NewEncoder(func(e *dynamodbattribute.Encoder) {
	//	e.EnableEmptyCollections = true
	//})

	av, errM := dynamodbattribute.MarshalMap(bodyMap)


	if errM != nil {
		return nil, errM
	}

	return av, nil

}

/**
* Convert the interface fields into a map
**/
func DecodeToMap(b io.ReadCloser, m interface{}) (interface{}, error) {

	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errJson := json.NewDecoder(b).Decode(&m)

	if errJson != nil {
		return nil, errJson
	}

	return m, nil
}

/**
* Model mapping of type interface to item from dynamodb
**/
func Unmarshal(result *dynamodb.GetItemOutput, m interface{}) map[string]interface{} {

	err := dynamodbattribute.UnmarshalMap(result.Item, &m)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	mapM, ok := m.(map[string]interface{})

	if !ok {
		fmt.Printf("ERROR: not a map-> %#v\n", m)
	}

	return mapM
}

func ParseEmptyCollection(av map[string]*dynamodb.AttributeValue, v string){

	av[v].NULL = nil
	av[v].M = map[string]*dynamodb.AttributeValue{}
}