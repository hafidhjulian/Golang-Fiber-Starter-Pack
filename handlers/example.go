package handlers

import (
	"golang-fiber-starterpack/models"

	"github.com/gofiber/fiber/v2"
)

func GetExample(c *fiber.Ctx) error {
	// data := c.FormValue("data")
	getdata := models.GetExample()
	return c.JSON(getdata)
}

func PostExample(c *fiber.Ctx) error {
	data := c.FormValue("data")
	getdata := models.PostExample(data)
	return c.JSON(getdata)
}
