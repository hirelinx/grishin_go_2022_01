package repositories

import (
	"controllers"
	"models"
)

type MessagesRepositoryReturns interface {
	bool | string | models.Message
}

type MessageId models.MessageId
type Message models.Message

type Messages interface {
	controllers.TokensManager
	CRUDRepository[MessageId, Message]
}

type messages struct {
}

func (m messages) Pass(strings []string) []any {
	panic("implement me")
}

func (m messages) Create(message *Message) (id MessageId) {
	{
		//TODO implement me
	}
	return
}

func (m messages) Read(id MessageId) (hasId Message, err error) {
	{
		//TODO implement me
	}
	return
}

func (m messages) Update(id MessageId, replacer Message) (old Message, err error) {
	{
		//TODO implement me
	}
	return
}

func (m messages) Delete(id MessageId) (removed Message, err error) {
	{
		//TODO implement me
	}
	return
}
