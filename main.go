//main.go

package main

func main(){
	a := App{}
	a.Initialize(
		//os.Getenv("APP_DB_USERNAME"),
		//os.Getenv("APP_DB_PASSWORD"),
		//os.Getenv("APP_DB_NAME"),
		"postgres", "1", "postgres",
	)

	a.Run(":8010")

	//fmt.Println("Hello world")
}
