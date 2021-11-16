package main


func main() {

	a := App{}
	a.InitializeConfig()
	a.InitializeDB()
	a.InitializeRouter()
	a.Run(":8080")
}
