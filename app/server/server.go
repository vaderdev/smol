package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vaderdev/smol/app/model"
	"github.com/vaderdev/smol/app/utils"
)

func redirect(ctx *fiber.Ctx) error {
	smolURL := ctx.Params("redirect")
	smol, err := model.FindBySmolUrl(smolURL)
	if err != nil {
		utils.SendMsg(ctx, fiber.StatusInternalServerError, "could not find smol link in db "+err.Error())
	}

	smol.Clicked += 1
	err = model.UpdateSmol(smol)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return ctx.Redirect(smol.Redirect, fiber.StatusTemporaryRedirect)
}

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

func createSmol(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var smol model.Smol
	err := ctx.BodyParser(&smol)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "error parsing JSON " + err.Error(),
		})
	}

	if smol.Random {
		smol.Smol = utils.RandomURL(8)
	}
	err = model.CreateSmol(smol)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "could not create smol in db " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(smol)
}

func updateSmol(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var smol model.Smol
	err := ctx.BodyParser(&smol)
	if err != nil {
		utils.SendMsg(ctx, fiber.StatusInternalServerError, "cannot parse json "+err.Error())
	}

	err = model.UpdateSmol(smol)
	if err != nil {
		utils.SendMsg(ctx, fiber.StatusInternalServerError, "could not update smol link in db "+err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(smol)
}

func deleteSmol(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		utils.SendMsg(ctx, fiber.StatusInternalServerError, "could not parse id from url "+err.Error())
	}
	err = model.DeleteSmol(id)
	if err != nil {
		utils.SendMsg(ctx, fiber.StatusInternalServerError, "could not delete smol link from db "+err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Deleted Smol Link",
	})
}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// redirect
	router.Get("/r/:redirect", redirect)

	// crud processes
	router.Get("/smol", getAllSmols)
	router.Get("/smol/:id", getSmol)
	router.Post("/smol", createSmol)
	router.Patch("/smol", updateSmol)
	router.Delete("/smol/:id", deleteSmol)

	router.Listen(":3000")
}
