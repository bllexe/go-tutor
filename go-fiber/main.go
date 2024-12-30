package main

import (
	"fmt"

	"github.com/bllexe/go-tutor/go-fiber/customer"
	"github.com/bllexe/go-tutor/go-fiber/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {

	app.Get("/api/customers", customer.GetCustomers)
	app.Post("/api/customer", customer.AddCustomer)
	app.Delete("/api/customer/:id", customer.DeleteCustomer)
	app.Get("/api/customer/:id", customer.GetCustomer)
}

func initDatabase() {
	database.ConnectDB()

	// Migrate the schema
	database.DB.AutoMigrate(&customer.Customer{})
	fmt.Println("Database migrated")
}

func main() {

	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(8080)

}
