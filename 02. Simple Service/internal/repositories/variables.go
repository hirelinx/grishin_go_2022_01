package repositories

import "controllers"

type Variables interface {
	controllers.TokensManager
}

type variables struct {
	Variables
}

//
//import "reflect"
//
//type Variables interface {
//	Get() (any2 *any, type2 reflect.Type)
//}
//
//type variables struct {
//	any2  *any
//	type2 reflect.Type
//}
//
//func (v variable) Get() (any2 *any, type2 reflect.Type) {
//	return v.any2, v.type2
//}
//
////func NewVariable(any2 *any, type2 reflect.Type) (v variable) {
////	if type2.
////}
//
//type ofVariables struct {
//}
//
////func (*ofVariables) Set(any2 *any, type2 reflect.Type) (err error){
////	any3 := any2
////	any2 = any3
////}
