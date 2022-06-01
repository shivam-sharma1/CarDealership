package service

import (
	"assignments/layeredcar/model"
	"github.com/google/uuid"
)

type Service interface {
	Create(c model.Car) (model.Car, error)
	Update(id uuid.UUID, c model.Car) (model.Car, error)
	Delete(id uuid.UUID) error
	GetById(id uuid.UUID) (model.Car, error)
	GetByBrand(brand string, engine string) ([]model.Car, error)
}
