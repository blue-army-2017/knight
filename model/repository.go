package model

import (
	"fmt"

	"gorm.io/gorm/clause"
)

type CRUDRepository[T Entity] interface {
	FindAll(orderBy string) ([]T, error)
	FindAllBy(key string, value any, orderBy string) ([]T, error)
	FindById(id string) (*T, error)
	Save(entity *T) error
	Delete(entity *T) error
}

type DefaultCRUDRepository[T Entity] struct {
}

func NewCRUDRepository[T Entity]() CRUDRepository[T] {
	return &DefaultCRUDRepository[T]{}
}

func (r *DefaultCRUDRepository[T]) FindAll(orderBy string) ([]T, error) {
	var entities []T
	result := db.
		Preload(clause.Associations).
		Order(orderBy).
		Find(&entities)
	return entities, result.Error
}

func (r *DefaultCRUDRepository[T]) FindAllBy(key string, value any, orderBy string) ([]T, error) {
	var entities []T
	result := db.
		Preload(clause.Associations).
		Where(fmt.Sprintf("%s = ?", key), value).
		Order(orderBy).
		Find(&entities)
	return entities, result.Error
}

func (r *DefaultCRUDRepository[T]) FindById(id string) (*T, error) {
	var entity T
	result := db.
		Preload(clause.Associations).
		First(&entity, WHERE_ID_IS, id)
	return &entity, result.Error
}

func (r *DefaultCRUDRepository[T]) Save(entity *T) error {
	result := db.Save(entity)
	return result.Error
}

func (r *DefaultCRUDRepository[T]) Delete(entity *T) error {
	result := db.Delete(entity)
	return result.Error
}
