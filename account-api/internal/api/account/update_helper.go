package account

import (
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
)

func UpdateValue(email string, cr *models.ChangeRequest) error{

	switch cr.Type {
	// string value
	case 1:
		return dynamodb.UpdateStringField(cr.Field,email,cr.NewString)
		break
	// map value
	case 2:
		return dynamodb.AppendNewMap(cr.NewMap.Uuid, email, &cr.NewMap, cr.Field)
		break
		// string value
	case 3:
		return dynamodb.UpdateBoolField(cr.Field,email,cr.NewBool)
		break
	}

	return nil
}
