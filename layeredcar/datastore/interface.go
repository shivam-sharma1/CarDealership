package datastore

import (
	"assignments/layeredcar/model"
	"github.com/google/uuid"
)

type Car interface {
	Create(car model.Car) (model.Car, error)
	GetById(id uuid.UUID) (model.Car, error)
	GetByBrand(brand string) ([]model.Car, error)
	Update(id uuid.UUID, car2 model.Car) (model.Car, error)
	Delete(id uuid.UUID) error
}

type Engine interface {
	Create(car model.Engine) (model.Engine, error)
	GetById(id uuid.UUID) (model.Engine, error)
	Update(id uuid.UUID, engine model.Engine) (model.Engine, error)
	Delete(id uuid.UUID) error
}
