package model

import (
	"strings"

	"gorm.io/gorm"
)

type Entity interface {
	Member | Season | SeasonGame
}

type CRUDRepository[T Entity] interface {
	FindAll(orderBy ...string) ([]T, error)
	FindById(id string) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type DefaultCRUDRepository[T Entity] struct {
}

func NewCRUDRepository[T Entity]() CRUDRepository[T] {
	return &DefaultCRUDRepository[T]{}
}

func (r *DefaultCRUDRepository[T]) FindAll(orderBy ...string) ([]T, error) {
	var entities []T
	var result *gorm.DB
	if len(orderBy) > 0 {
		result = db.
			Order(strings.Join(orderBy, ",")).
			Find(&entities)
	} else {
		result = db.
			Order("id").
			Find(&entities)
	}
	return entities, result.Error
}

func (r *DefaultCRUDRepository[T]) FindById(id string) (*T, error) {
	var entity T
	result := db.First(&entity, "id = ?", id)
	return &entity, result.Error
}

func (r *DefaultCRUDRepository[T]) Create(entity *T) error {
	result := db.Create(entity)
	return result.Error
}

func (r *DefaultCRUDRepository[T]) Update(entity *T) error {
	result := db.Save(entity)
	return result.Error
}

func (r *DefaultCRUDRepository[T]) Delete(entity *T) error {
	result := db.Delete(entity)
	return result.Error
}
