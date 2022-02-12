package lead

import(
	"github.com/FuaAlfu/golang-multi-projects/tree/master/simple-crm/database"
	
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name      string
	Company   string
	Email     string
	phone     int
}

func GetLead(){}

func GetLeads(){}

func NewLead(){}

func DeleteLead(){}