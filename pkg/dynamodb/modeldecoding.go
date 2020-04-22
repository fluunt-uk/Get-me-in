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
<<<<<<< HEAD

}

// Have to change this so it can accommodate all types of structs, need to pass in interface as param
func ConvertStructToDynamoAttribute(sub interface{}) (map[string]*dynamodb.AttributeValue, error) {

	av, errM := dynamodbattribute.MarshalMap(sub)

	if errM != nil {
		return nil, errM
	}

	return av, nil
}

/**
* Convert type interface to dynamodb readable object and JSON
**/
func DecodeToDynamoAttributeAndJson(readBody io.ReadCloser, m interface{}) (map[string]*dynamodb.AttributeValue, error, string) {

	bodyMap, err := DecodeToMap(readBody, m)
	jsonString, err := json.Marshal(bodyMap)

	if err != nil {
		return nil, err, ""
	}

	av, errM := dynamodbattribute.MarshalMap(bodyMap)

	if errM != nil {
		return nil, errM, ""
	}

	return av, nil, string(jsonString)

=======
>>>>>>> ef8155820d93a627b42eccadab034a303e096200
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

func ParseEmptyCollection(av map[string]*dynamodb.AttributeValue, v string){

	av[v].NULL = nil
	av[v].M = map[string]*dynamodb.AttributeValue{}
}