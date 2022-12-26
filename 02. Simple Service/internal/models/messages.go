package models

type Message struct {
	Identificator[MessageId]
	DestUser UserId
	SendUser UserId
	Text     string
}
type MessageId = uint64

func (m Message) Identify() MessageId {
	return m.Identificator.id()
}

func (m Message) Reidentify(new MessageId) {
	m.Identificator.setId(new)
}

func NewMessage(identificator Identificator[MessageId], destUser UserId, sendUser UserId, text string) (*Message, error) {
	return &Message{Identificator: identificator, DestUser: destUser, SendUser: sendUser, Text: text}, nil
}
