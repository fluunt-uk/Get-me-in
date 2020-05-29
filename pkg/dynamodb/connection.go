package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"net/http"
)

/*** Values injected from main internal that imports this library ***/
type Wrapper struct {
	Connection *dynamodb.DynamoDB
	SearchParam *string
	GenericModel *interface{}
	Table *string
	Credentials *credentials.Credentials
	Region *string
}
/*******************************************************************/

//Create a connection to DB and assign the session to our struct variable
//connection variable is shared by other repo(CRUD)
func (d *Wrapper) DefaultConnect() error {

	sess, err := newSession(d.Table, d.SearchParam, d.GenericModel, d.Region, d.Credentials)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	//creating the actual session
	d.Connection = dynamodb.New(sess)

	return nil
}

//SOLID
//O: Open for Modification - might come in handy
func (d *Wrapper) CustomConnect(t *string, s *string, gm *interface{}, r *string, c *credentials.Credentials) (*dynamodb.DynamoDB,error) {

	sess, err := newSession(t, s, gm, r, c)

	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}

func newSession(t *string, s *string, gm *interface{}, r *string, c *credentials.Credentials) (*session.Session, error){
	//defensive coding, checking for empty values
	if *t  == "" && *s == "" && *gm == nil {
		return nil, &ErrorString{
			Reason: "Injected values are empty or nil",
			Code:   http.StatusBadRequest,
		}
	}

	//creating the object
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(*r),
		Credentials: c,
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
