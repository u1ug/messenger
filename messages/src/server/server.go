package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fastjson"
	"messenger/config"
	"messenger/pkg/database"
	"messenger/pkg/logging"
)

var (
	DB     *sql.DB
	logger = logging.GetLogger()
	parser fastjson.Parser
)

func main() {
	var err error
	DB, err = database.NewDefaultConnection("root:123456@/messenger")
	if err != nil {
		logger.Panicln(err)
	}
	logger.Infoln("db connection pool created")

	app := fiber.New()
	app.Use(cors.New())
	initRoutes(app)

	logger.Infoln("running server at ", config.MessagesServicePort)
	err = app.Listen(config.MessagesServicePort)
	if err != nil {
		logger.Panicln(err)
	}
}
