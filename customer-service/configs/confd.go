package configs


//name string, intro string, instruc string, buttonText string, buttonColor string, buttonLink string, outro string
type ActionEmail struct {
	Name 		string
	Intro 		string
	Instruct 	string
	ButtonText 	string
	ButtonColor string
	ButtonLink 	string
	Outro 		string
}

type NotificationEmail struct {
	Name 		string
	Intro 		string
	Outro 		string
}