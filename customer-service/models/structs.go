package models

//name string, intro string, instruc string, buttonText string, buttonColor string, buttonLink string, outro string
type ActionEmailStruct struct {
	Name 		string
	Intro 		string
	Instruct 	string
	ButtonText 	string
	ButtonColor string
	ButtonLink 	string
	Outro 		string
}

type NotificationEmailStruct struct {
	Name  string
	Intro string
	Outro string
}