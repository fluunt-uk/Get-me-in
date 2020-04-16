package account

import (
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
)

func UpdateValue(email string, cr *models.ChangeRequest){

	switch cr.Type {
	// string value
	case 1:
		dynamodb.UpdateSingleField(cr.Field,email,cr.NewString)
		break
	// map value
	case 2:
		dynamodb.AppendNewMap(cr.NewMap.Uuid, email, &cr.NewMap, cr.Field)
	}

}
