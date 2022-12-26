package repositories

import (
	"controllers"
	"models"
)

type UserRepositoryReturns interface {
	bool | string | models.User
}

type Users interface {
	CRUDRepository[uint32, models.User]
	controllers.TokensManager
}

type users struct {
}

func (u users) Pass(strings []string) []any {
	//TODO implement me
	panic("implement me")
}
