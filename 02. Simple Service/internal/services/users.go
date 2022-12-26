package services

import (
	"controllers"
)

type Users interface {
	controllers.TokensManager
}

type users struct {
	Users
}
