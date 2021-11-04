package main

var (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "beerdb"
	dbHost     = "localhost"
	dbPort     = "5432"
	sslMode    = "disable"
)

func main() {

	a := App{}

	a.InitializeDB(dbUser, dbPassword, dbName, dbHost, dbPort, sslMode)
	a.InitializeRouter()
	a.Run(":8080")
}
