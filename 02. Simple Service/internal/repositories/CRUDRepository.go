package repositories

import "models"

type CRUDRepository[Id any, V models.HasId[Id]] interface {
	Create(v V) (id Id)
	Read(id Id) (hasId V, err error)
	Update(id Id, replacer V) (old V, err error)
	Delete(id Id) (removed V, err error)
}
