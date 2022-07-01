package entities

type Message struct {
	MessageID    string `json:"message_id"`
	SenderID     int64  `json:"sender_id"`
	PeerID       int64  `json:"peer_id"`
	ReplyID      string `json:"reply_id"`
	Text         string `json:"text"`
	AttachmentID int64  `json:"attachment"`
	Timestamp    int64  `json:"timestamp"`
}

func NewMessage(sender int64, receiver int64, replyID string, text string, attachment int64) *Message {
	return &Message{
		SenderID:     sender,
		PeerID:       receiver,
		ReplyID:      replyID,
		Text:         text,
		AttachmentID: attachment,
	}
}

func (m *Message) IsValid() bool {
	if m.SenderID == 0 || m.PeerID == 0 {
		return false
	}
	if m.Text == "" {
		return false
	}
	return true
}
