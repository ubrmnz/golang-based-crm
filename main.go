package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thecodemensch/golang-based-crm/database"
)

func route_setups(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLeads)
	app.Delete(DeleteLeads)
}

func init_database() {

}

func main() {
	app := fiber.New()
	init_database()
	route_setups(app)
	app.Listen("3000")
	defer database.DBConn.Close()
}
