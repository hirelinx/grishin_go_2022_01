package controllers

type Users interface {
	TokensManager
}

type users struct {
	Users
}
