package controllers

type TokensManager interface {
	Pass([]string) []any
}
