package controllers

import "services"

type Messages interface {
	TokensManager
}

type messages struct {
	Messages
	service *services.Messages
}
