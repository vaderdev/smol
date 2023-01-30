package utils

import "github.com/gofiber/fiber/v2"

func SendMsg(ctx *fiber.Ctx, status int, msg string) {
	ctx.Status(status).JSON(fiber.Map{
		"msg": msg,
	})
}
