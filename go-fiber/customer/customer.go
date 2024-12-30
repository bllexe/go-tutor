package customer

import (
	"github.com/bllexe/go-tutor/go-fiber/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

func GetCustomers(c *fiber.Ctx) {
	//get all customers form db
	db := database.DB
	var customers []Customer
	if err := db.Find(&customers).Error; err != nil {
		c.SendStatus(400)
	}
	c.JSON(customers)
}

func AddCustomer(c *fiber.Ctx) {
	db := database.DB
	customer := new(Customer)
	if err := c.BodyParser(customer); err != nil {
		c.SendStatus(400)
	}
	db.Create(&customer)
	c.JSON(customer)
}

func GetCustomer(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var customer Customer
	db.Find(&customer, id)
	c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var customer Customer
	db.First(&customer, id)
	if customer.ID == 0 {
		c.Status(404).Send("Customer not found")
	}
	db.Delete(&customer, id)
}
