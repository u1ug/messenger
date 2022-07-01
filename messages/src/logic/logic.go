package logic

import (
	"database/sql"
	"github.com/google/uuid"
	"messenger/messages/src/entities"
	"messenger/pkg/logging"
	"time"
)

var logger = logging.GetLogger()

func InsertMessage(db *sql.DB, msg *entities.Message) (error, *entities.Message) {
	msg.MessageID = uuid.New().String()
	msg.Timestamp = time.Now().UnixMicro()

	_, err := db.Exec("insert into messages(MessageID, SenderID, PeerID, ReplyID, Text, AttachmentID, Timestamp) values(?,?,?,?,?,?,?)", msg.MessageID, msg.SenderID, msg.PeerID, msg.ReplyID, msg.Text, msg.AttachmentID, msg.Timestamp)
	if err != nil {
		logger.Errorln(err)
	}
	return err, msg
}
