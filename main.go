package main


var (
	dbUser = "postgres"
	dbPassword = "postgres"
	dbName = "beerdb"
	dbHost = "localhost"
	dbPort = "5432"
)
func main() {

	a := App{}

	a.Initialize(dbUser,dbPassword,dbName,dbHost,dbPort)
	a.Run(":8080")
}
