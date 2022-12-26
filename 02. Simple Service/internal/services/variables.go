package services

import "controllers"

type Variables interface {
	controllers.TokensManager
}

type variables struct {
	Variables
}
