package controllers

import "github.com/gofiber/fiber/v2"

func (ctr *Controllers) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"Message": "SUCCESS",
	})
}
