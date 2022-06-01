package service

import (
	"assignments/layeredcar/datastore"
	"assignments/layeredcar/model"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

type service struct {
	carStore    datastore.Car
	engineStore datastore.Engine
}

func New(store datastore.Car, store2 datastore.Engine) service {
	return service{carStore: store, engineStore: store2}
}

func (s service) Create(c model.Car) (model.Car, error) {

	//Empty entry case
	k := model.Car{Id: uuid.Nil, Name: "", Year: 0, Brand: "", FuelType: "", Engine: model.Engine{EngineId: uuid.Nil, Displacement: 0, Noc: 0, Rng: 0}}
	if c == k {
		return k, errors.New("no entry created, Empty data sent")
	}

	//Invalid year case
	if c.Year < 1980 || c.Year > time.Now().Year() {
		return c, errors.New("No entry created,Invalid Year")
	}

	//Invalid brand case
	brand := []string{"Tesla", "BMW", "Ferrari", "Porsche", "Mercedes"}
	f := 0
	for _, b := range brand {
		if c.Brand == b {
			f = 1
			break
		}
	}
	if f == 0 {
		return c, errors.New("no entry created,Invalid Brand")
	}
	e, err := s.engineStore.Create(c.Engine)
	if err != nil {
		return model.Car{}, err
	}

	c.Engine.EngineId = e.EngineId

	car, err := s.carStore.Create(c)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}

func (s service) Update(id uuid.UUID, c model.Car) (model.Car, error) {

	k := model.Car{Id: uuid.Nil, Name: "", Year: 0, Brand: "", FuelType: "", Engine: model.Engine{EngineId: uuid.Nil, Displacement: 0, Noc: 0, Rng: 0}}
	if k == c {
		return k, errors.New("no entry updated, Empty data sent")
	}
	//Invalid year case
	if c.Year < 1980 || c.Year > time.Now().Year() {
		return c, errors.New("no entry updated,Invalid Year")
	}

	//Invalid brand case
	brand := []string{"Tesla", "BMW", "Ferrari", "Porsche", "Mercedes"}
	f := 0
	for _, b := range brand {
		if c.Brand == b {
			f = 1
			break
		}
	}
	if f == 0 {
		return c, errors.New("no entry updated,Invalid Brand")
	}

	car, err := s.carStore.Update(id, c)
	if err != nil {
		return model.Car{}, err
	}
	cr, err := s.carStore.GetById(id)
	if err != nil {
		log.Println("Error getting id for update")
		return model.Car{}, err
	}
	e, err := s.engineStore.Update(cr.Engine.EngineId, c.Engine)
	if err != nil {
		log.Println("Error updating engine")
		return model.Car{}, err
	}
	car.Engine = e
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (s service) Delete(id uuid.UUID) error {
	var c model.Car
	if id == uuid.Nil {
		return errors.New("cannot delete,empty id")
	}
	c, err := s.carStore.GetById(id)
	if err != nil {
		return errors.New("Error getting id while deleting")
	}
	err = s.carStore.Delete(id)
	if err != nil {
		return errors.New("Error while deleting car")
	}
	err = s.engineStore.Delete(c.Engine.EngineId)
	if err != nil {
		return errors.New("Error while deleting engine")
	}
	return nil
}

func (s service) GetById(id uuid.UUID) (model.Car, error) {
	var c model.Car
	var e model.Engine
	if id == uuid.Nil {
		return model.Car{}, errors.New("cannot Get,empty id")
	}
	c, err := s.carStore.GetById(id)
	if err != nil {
		return model.Car{}, errors.New("Error getting data by id")
	}
	e, err = s.engineStore.GetById(c.Engine.EngineId)
	if err != nil {
		return model.Car{}, errors.New("Error while getting engine")
	}
	c.Engine = e
	return c, nil
}

func (s service) GetByBrand(brand string, engine string) ([]model.Car, error) {
	var c []model.Car
	var e model.Engine

	if engine == "" || brand == "" {
		return []model.Car{}, errors.New("can not get,empty brand")
	}
	c, err := s.carStore.GetByBrand(brand)
	if err != nil {
		return []model.Car{}, errors.New("Error getting data by brand")
	}

	if engine == "included" {
		for i, _ := range c {
			e, err = s.engineStore.GetById(c[i].Engine.EngineId)
			if err != nil {
				return []model.Car{}, errors.New("Error while getting engine")
			}
			c[i].Engine = e
		}
	}
	return c, nil
}
