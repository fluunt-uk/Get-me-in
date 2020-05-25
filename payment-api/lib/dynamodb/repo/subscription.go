package repo

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Wrapper struct {
	//dynamo client
	DC		*dynamodb.Wrapper
}
//implement only the necessary methods for each repository
//available to be consumed by the API
type Builder interface{
	Create(*models.Subscription) (string, error)
	Del(string) (string, error)
}

//get all the adverts for a specific account
//token validated
func (s *Wrapper) Create(body *models.Subscription) (string, error) {

	av, errM := dynamodbattribute.MarshalMap(body)

	if errM != nil {
		return "", errM
	}

	fmt.Println(av)
	err := s.DC.CreateItem(av)

	if err != nil {
		// Need to handle changing premium status here, will need to call endpoint
		return "Failed", err
	}
	return "Success", err
}

func (s *Wrapper) Del(email string) (string, error) {
	err := s.DC.DeleteItem(email)

	if err != nil {
		return "Failed to delete item", err
	}
	return "Item deleted", nil
}

