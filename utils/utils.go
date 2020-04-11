//internal specific configs are loaded at runtime
func loadEnvConfigs(name &string, port &int,searchParam &string ) {

	log.Println("Running on %s \n", port)
	dynamodb.SearchParam = configs.UNIQUE_IDENTIFIER

	switch env := os.Getenv("ENV"); env {
	case "UAT":
		dynamodb.DynamoTable = "uat-" + name
	case "PROD":
		dynamodb.DynamoTable = "prod-" + name
	default:
		dynamodb.DynamoTable = "dev-users"
	}

	log.Println("Environment:" + env)
}