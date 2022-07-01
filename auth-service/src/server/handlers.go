package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"messenger/auth-service/src/logic"
)

func initRoutes(app *fiber.App) {
	app.Post("/auth/get_sha256", handleGetSHA256)
}

func handleGetSHA256(c *fiber.Ctx) error {
	v, err := parser.ParseBytes(c.Body())
	if err != nil {
		logger.Errorln(err)
		return errors.New("{\"error\":\"invalid request\"}")
	}
	userID := v.GetInt64("user_id")
	hash, err := logic.GetUserSHA256(DB, userID)
	if err != nil {
		logger.Errorln(err)
	}
	err = c.SendString(hash)
	if err != nil {
		logger.Errorln(err)
	}
	return err
}
