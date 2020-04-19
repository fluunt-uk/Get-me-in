package subscription

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/payment-api/lib/stripe-api/resources/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
)

func AddSubscription(body models.Subscription) (string, error) {

	dynamoAttr, errDecode := dynamodb.ConvertStructToDynamoAttribute(body)

	if errDecode != nil {
		return "Error: ", errDecode
	}

	fmt.Println(dynamoAttr)
	err := dynamodb.CreateItem(dynamoAttr)

	if err != nil {
		// Need to handle changing premium status here, will need to call endpoint
		return "Failed", err
	}
	return "Success", err
}

// Will need to get email from somewhere, not sure where yet
func DeleteSubscription(email string) (string, error) {
	err := dynamodb.DeleteItem(email)

	if err != nil {
		return "Failed to delete item", err
	}
	return "Item deleted", nil
}

func UpdateSubscription(){}
