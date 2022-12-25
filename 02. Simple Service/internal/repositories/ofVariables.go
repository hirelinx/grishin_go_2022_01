package repositories

import "reflect"

type Variable interface {
	Get() (any2 *any, type2 reflect.Type)
}

type variable struct {
	any2  *any
	type2 reflect.Type
}

func (v variable) Get() (any2 *any, type2 reflect.Type) {
	return v.any2, v.type2
}

//func NewVariable(any2 *any, type2 reflect.Type) (v variable) {
//	if type2.
//}

type ofVariables struct {
}

//func (*ofVariables) Set(any2 *any, type2 reflect.Type) (err error){
//	any3 := any2
//	any2 = any3
//}
