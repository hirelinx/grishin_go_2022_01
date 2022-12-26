package services

import "controllers"

type Messages interface {
	controllers.TokensManager
}

type messages struct {
	Messages
}
