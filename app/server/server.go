package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vaderdev/smol/app/model"
)

func getAllSmols(ctx *fiber.Ctx) error {
	smols, err := model.GetAllSmols()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all smol links " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(smols)
}

func getSmol(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "error parsing the id " + err.Error(),
		})
	}

	smol, err := model.GetSmol(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "error could not retrieve smol link from db " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(smol)
}

func createGoly(ctx *fiber.Ctx) error {

}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/smol", getAllSmols)
	router.Get("/smol/:id", getSmol)

	router.Listen(":3000")
}
