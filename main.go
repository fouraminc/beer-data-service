package main

var (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "beerdb"
	dbHost     = "localhost"
	dbPort     = "5432"
)

func main() {

	a := App{}

	a.InitializeDB(dbUser, dbPassword, dbName, dbHost, dbPort)
	a.InitializeRouter()
	a.Run(":8080")
}
