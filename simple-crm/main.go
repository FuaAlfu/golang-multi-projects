package main

import(
	"fmt"

	"github.com/FuaAlfu/golang-multi-projects/tree/master/simple-crm/database"
	"github.com/FuaAlfu/golang-multi-projects/tree/master/simple-crm/lead"
	"github.com/gofiber/fiber"
)

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


func setupRoutes(app *fiber.App){
	app.GET("/", homePage)
	app.GET("/api/leads",lead.GetLeads)
	app.GET("/api/leads/:id",lead.GetLead)
	app.POST(lead.NewLead)
	app.DELETE(lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("connection opend to database..")
	database.DBConn.AutoMigrate(&lead.Lead())
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	initDatabase()
	app.Listen(3000)

	defer database.DBConn.Close()
}