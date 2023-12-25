package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazimovzaman2/online_book/pkg/utils"
)

func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
