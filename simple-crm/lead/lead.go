package lead

import(
	"github.com/FuaAlfu/golang-multi-projects/tree/master/simple-crm/database"
	
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name      string   `json: "name"`
	Company   string   `json: "company"`
	Email     string   `json: "email"`
	phone     int      `json: "phone"`
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.Find(&lead,id) //instance of connection
	c.json(lead)
}

func GetLeads(c *fiber.Ctx){
	db := database.DBConn
	var leads []Lead

	db.Find(&leads)
	c.json(leads)
}

func NewLead(c *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err == nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.json(lead)
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead,id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found with id")
		return
	}
	db.Delete(&lead)
	c.Send("lead deleted..")
}