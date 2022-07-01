package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fastjson"
	"messenger/config"
	"messenger/pkg/logging"
)

var (
	logger = logging.GetLogger()
	parser fastjson.Parser
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initRoutes(app)

	logger.Infoln("running server at ", config.ApiRouterServicePort)
	err := app.Listen(config.ApiRouterServicePort)
	if err != nil {
		logger.Panicln(err)
	}
}
