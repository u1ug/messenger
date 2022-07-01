package main

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"messenger/messages/src/entities"
	"messenger/messages/src/logic"
)

func handleSendMessage(c *fiber.Ctx) error {
	v, err := parser.ParseBytes(c.Body())
	if err != nil {
		return errors.New("{\"error\":\"invalid request\"}")
	}
	senderID := v.GetInt64("sender_id")
	receiverID := v.GetInt64("peer_id")
	replyID := string(v.GetStringBytes("reply_id"))
	text := string(v.GetStringBytes("text"))
	attachment := v.GetInt64("attachment")

	msg := entities.NewMessage(senderID, receiverID, replyID, text, attachment)
	if !msg.IsValid() {
		logger.Errorln(err)
		return errors.New("{\"error\":\"invalid request\"}")
	}
	err, msg = logic.InsertMessage(DB, msg)
	if err != nil {
		logger.Errorln(err)
		return errors.New("{\"error\":\"internal server error\"}")
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		logger.Errorln(err)
	}
	err = c.Send(jsonData)
	if err != nil {
		logger.Errorln(err)
	}
	return nil
}

func initRoutes(app *fiber.App) {
	app.Post("/messages/send_message", handleSendMessage)
}
