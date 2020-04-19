package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

/*** Values injected from main internal that imports this library ***/
type DynamoDB struct {
	Connection *dynamodb.DynamoDB
	SearchParam string
	GenericModel interface{}
	Table string
	Credentials *credentials.Credentials
	Region string
}
/*******************************************************************/

//Create a connection to DB and assign the session to DynamoConnection variable
//DynamoConnection variable is shared by other repo(CRUD)
func (d *DynamoDB) Connect() error {

	//defensive coding, checking for empty values
	if d.Table == "" && d.SearchParam == "" && d.GenericModel == nil {
		return &ErrorString{
			Reason: "Injected values are empty or nil",
			Code:   http.StatusBadRequest,
		}
	}

	//creating the object
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(d.Region),
		Credentials: d.Credentials,
	})

	if err != nil {
		return err
	}

	//creating the actual session
	d.Connection = dynamodb.New(sess)

	return nil
}
