package main

import(
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type(
	MongoInstance struct{
		Clint
		DB
	}
	Employee struct{
		ID     string  `json: "id"`
		Name   string  `json: "name"`
		Salary float64 `json: "salary"`
		Age    float64 `json: "age"`
	}
)

var mg MongoInstance

const(
	dbName   =   "fiber-hrms"
	mongoURI =   "mongodb://localhost:27017" + dbName  //1
)

func connect()error{

}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, `<!DOCTYPE html>
	<html lang="en">
	
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="style.css">
		<title>my page</title>
	</head>
	
	<body>
		<form class="form" action="/ouath/github" method="post">
			<input type="submit" value="login with github">
		</form>
	</body>

	<style>
		 *{
			 box-border: box-sizing;
			 background-color: #ccc;
		 }
		 .form{
		   background-color: #eee;
		 }
	</style>
	
	</html>`)
}

func getEmployees(c *fiber.Ctx)error{

}

func getEmployee(c *fiber.Ctx)error{
	
}

func newEmployee(c *fiber.Ctx)error{
	
}

func deleteEmployee(c *fiber.Ctx)error{
	
}


func setupRoutes(app *fiber.App){
	app.GET("/", homePage)
	app.GET("/employee/:id",getEmployees)
	app.GET("/employee/",getEmployee)
	app.POST("/employee",newEmployee)
	app.DELETE("/employee/:id"deleteEmployee)
}

func main() {
	fmt.Println("start")
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}