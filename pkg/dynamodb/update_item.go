package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io"
	"log"
)

/**
TODO: not working as expected yet
***** NOT IN USE, under implementation ****
**/
func UpdateSingleField(fieldToUpdate string, recordToUpdate string, newValue string) (bool, error) {

	input := &dynamodb.UpdateItemInput{
		//values that needs change
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":" + fieldToUpdate: {
				S: aws.String(newValue),
			},
		},
		TableName: aws.String(DynamoTable),
		//search by(primary key)
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(recordToUpdate),
			},
		},
		// May need updating
		ReturnValues:     aws.String("UPDATED_NEW"),
		//set {field name on dynamoDb} = {fieldtoUpdate}
		UpdateExpression: aws.String("set " + fieldToUpdate + "= :" + fieldToUpdate),
	}

	_, err := DynamoConnection.UpdateItem(input)
	if err != nil {
		return false, err
	}

	return true, nil
}


//func AppendNewObject(){
//
//	//value, _ := GetItem("luno@gmail.com")
//
//	////advert
//	advertDetails := map[string]*dynamodb.AttributeValue{
//		"advert_id": {
//			S: aws.String("asd"),
//		},
//		"advert_description": {
//			S: aws.String("what, just trying"),
//		},
//	}
//
//
//	//application object
//	newAdvert := 	map[string]*dynamodb.AttributeValue{
//		"advert_id81273984" :{
//			M: advertDetails,
//		},
//	}
//
//
//
//	av := &dynamodb.AttributeValue{
//		S: aws.String("changed it"),
//	}
//	var qids []*dynamodb.AttributeValue
//	qids = append(qids, av)
//
//
//	input := &dynamodb.UpdateItemInput{
//		Key: map[string]*dynamodb.AttributeValue{
//			"email": {
//				S: aws.String("lunos@gmail.com"),
//			},
//		},
//		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
//			":applications": {
//				L: qids,
//			},
//			":empty_list": {
//				L: []*dynamodb.AttributeValue{},
//			},
//		},
//		ReturnValues:     aws.String("ALL_NEW"),
//		UpdateExpression: aws.String("SET applications = list_append(if_not_exists(applications, :empty_list), :applications)"),
//		TableName:        aws.String("dev-users"),
//	}
//	_, err := DynamoConnection.UpdateItem(input)
//	if err != nil {
//		log.Println(err)
//	}
//}

func AppendNewMap(id string, r string, body io.ReadCloser, i interface{}){

	//advertDetails := map[string]*dynamodb.AttributeValue{
	//	"advert_id": {
	//		S: aws.String("asd"),
	//	},
	//	"advert_description": {
	//		S: aws.String("what, just trying"),
	//	},
	//}

	m, _ := DecodeToDynamoAttribute(body, i)

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(r),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":applications": {
				M: m,
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET applications." + id + " = :applications"),
		TableName:        aws.String("dev-users"),
	}
	_, err := DynamoConnection.UpdateItem(input)
	if err != nil {
		log.Println(err)
	}

}