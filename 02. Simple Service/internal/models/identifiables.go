package models

//type Getter[T any] interface {
//	Get() T
//}
//
//type MyStruct[T any] struct {
//	Val T
//}
//
//func (m MyStruct[T]) Get() T {
//	return m.Val
//}
//
//func bar[T any]() Getter[T] {
//	return MyStruct[T]{} // ok
//}

type HasId[Id any] interface {
	Identify() Id
	Reidentify(new Id)
}

type Identificator[Id any] interface {
	id() Id
	setId(id Id)
}

type IdentificatorImpl[Id any] struct {
	id Id
}

func NewIdentificator[Id any](id Id) *IdentificatorImpl[Id] {
	return &IdentificatorImpl[Id]{id: id}
}

func (h IdentificatorImpl[Id]) Id() Id {
	return h.id
}

func (h IdentificatorImpl[Id]) SetId(id Id) {
	h.id = id
}
