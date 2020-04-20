package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (d *Wrapper) UpdateStringField(fieldToUpdate string, recordToUpdate string, newValue string) error {

	input := &dynamodb.UpdateItemInput{
		//values that needs change
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":" + fieldToUpdate: {
				S: aws.String(newValue),
			},
		},
		TableName: d.Table,
		//search by(primary key)
		Key: map[string]*dynamodb.AttributeValue{
			*d.SearchParam: {
				S: aws.String(recordToUpdate),
			},
		},
		// Not used at the moment
		//ReturnValues:     aws.String("UPDATED_NEW"),
		//set {field name on dynamoDb} = {fieldtoUpdate}
		UpdateExpression: aws.String("set " + fieldToUpdate + "= :" + fieldToUpdate),
	}

	_, err := d.Connection.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (d *Wrapper) UpdateBoolField(fieldToUpdate string, recordToUpdate string, newValue bool) error{

	input := &dynamodb.UpdateItemInput{
		//values that needs change
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":" + fieldToUpdate: {
				BOOL: aws.Bool(newValue),
			},
		},
		TableName: d.Table,
		//search by(primary key)
		Key: map[string]*dynamodb.AttributeValue{
			*d.SearchParam: {
				S: aws.String(recordToUpdate),
			},
		},
		// Not used at the moment
		//ReturnValues:     aws.String("UPDATED_NEW"),
		//set {field name on dynamoDb} = {fieldtoUpdate}
		UpdateExpression: aws.String("set " + fieldToUpdate + "= :" + fieldToUpdate),
	}

	_, err := d.Connection.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil


}

func (d *Wrapper) AppendNewMap(mapId string, r string, i interface{}, key string) error {

	m, _ := dynamodbattribute.MarshalMap(&i)

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			*d.SearchParam: {
				S: aws.String(r),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":" + key: {
				M: m,
			},
		},
		//Not used at the moment
		//ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET "+ key+ "." + mapId + " = :" + key),
		TableName:        d.Table,
	}
	_, err := d.Connection.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}

//TODO:append to lists
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