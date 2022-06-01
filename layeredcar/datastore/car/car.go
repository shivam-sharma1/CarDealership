package car

import (
	"assignments/layeredcar/model"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type carstorer struct {
	db *sql.DB
}

func New(Db *sql.DB) carstorer {
	return carstorer{db: Db}
}

func (c carstorer) GetById(id uuid.UUID) (model.Car, error) {
	var car2 model.Car
	res := c.db.QueryRow(SelectByIdQuery, id)
	var cid string
	err := res.Scan(&cid, &car2.Name, &car2.Year, &car2.Brand, &car2.FuelType, &car2.Engine.EngineId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Car{}, err
		}
		return model.Car{}, err
	}

	car2.Id, err = uuid.Parse(cid)
	if err != nil {
		return model.Car{}, err
	}
	return car2, nil
}
func (c carstorer) GetByBrand(brand string) ([]model.Car, error) {
	var res []model.Car
	rows, err := c.db.Query(SelectByBrandQuery, brand)
	if err != nil {
		return []model.Car{}, err
	}

	for rows.Next() {
		var r model.Car
		err := rows.Scan(&r.Id, &r.Name, &r.Year, &r.Brand, &r.FuelType, &r.Engine.EngineId)
		if err != nil {
			if err == sql.ErrNoRows {
				return []model.Car{}, err
			}
			return []model.Car{}, err
		}
		res = append(res, r)
	}

	return res, nil
}
func (c carstorer) Create(car model.Car) (model.Car, error) {
	id := uuid.New()
	car.Id = id
	_, err := c.db.Exec(InsertQuery, car.Id.String(), car.Name, car.Year, car.Brand, car.FuelType, car.Engine.EngineId)
	if err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (c carstorer) Update(id uuid.UUID, car2 model.Car) (model.Car, error) {
	_, err := c.db.Exec(UpdateQuery, car2.Name, car2.Year, car2.Brand, car2.FuelType, id.String())
	if err != nil {
		return model.Car{}, err
	}
	return car2, nil
}

func (c carstorer) Delete(id uuid.UUID) error {
	_, err := c.db.Exec(DeleteQuery, id.String())
	if err != nil {
		log.Println("Error while deleting!", err)
		return err
	}

	return nil
}

//First try
//package car
//
//import (
//	"assignments/layeredcar/models"
//	"context"
//	"database/sql"
//	"errors"
//	"log"
//)
//
//type CarStorer struct {
//	db *sql.DB
//}
//
//func New(db *sql.DB) CarStorer {
//	return CarStorer{db}
//}
//
//func (c CarStorer) CarStorer_Post(ctx context.Context, car *models.Car) (models.Car, error) {
//	var err error
//	//car.Id = uuid.New()
//	_, err = c.db.ExecContext(ctx, "insert into cars values(?,?,?,?,?,?)", car.Id, car.Engine.Engineid, car.Name, car.Year, car.Brand, car.FuelType)
//	if err != nil {
//		return models.Car{}, errors.New("no entry created")
//	}
//	//_, err = c.db.Exec("insert into engine values(?,?,?,?,?)", uuid.New(), car.Engine.Displacement, car.Engine.Noc, car.Engine.Rnge, car.Id)
//	//if err != nil {
//	//	return models.Car{}, err
//	//}
//	return *car, nil
//}
//func (c CarStorer) CarStorer_Put(ctx context.Context, id string, car *models.Car) (models.Car, error) {
//	var err error
//	//car.Id = uuid.New()
//	_, err = c.db.ExecContext(ctx, "update cars set name=?,year=?,brand=?,fueltype=? where id=?", car.Name, car.Year, car.Brand, car.FuelType, id)
//	if err != nil {
//		return models.Car{}, errors.New("no new values")
//	}
//	return *car, nil
//}
//
//func (c CarStorer) CarStorer_GetID(ctx context.Context, id string) (models.Car, error) {
//	var car models.Car
//	err := c.db.QueryRowContext(ctx, "select * from cars where id=?;", id).
//		Scan(&car.Id, &car.Engine.Engineid, &car.Name, &car.Year, &car.Brand, &car.FuelType)
//	if err != nil {
//		return models.Car{}, err
//	}
//	return car, nil
//}
//
//func (c CarStorer) GetCarsByBrand(ctx context.Context, brand string) ([]models.Car, error) {
//	var car []models.Car
//
//	result, err := c.db.QueryContext(ctx, "select * from Car where brand=?;", brand)
//	if err != nil {
//		return nil, err
//	}
//
//	defer func() {
//		err = result.Err()
//		if err != nil {
//			log.Printf("error: %v", err)
//		}
//	}()
//
//	defer func() {
//		err = result.Close()
//		if err != nil {
//			log.Printf("error: %v", err)
//		}
//	}()
//	for result.Next() {
//		var s models.Car
//		err = result.Scan(&s.Id, &s.Engine.Engineid, &s.Name, &s.Year, &s.Brand, &s.FuelType)
//		if err != nil {
//			return nil, err
//		}
//		car = append(car, s)
//	}
//	return car, nil
//}
//
//func (c CarStorer) DeleteCar(ctx context.Context, id string) (models.Car, error) {
//	_, err := c.db.ExecContext(ctx, "delete from cars where id=?", id)
//	if err != nil {
//		return models.Car{}, err
//	}
//	return models.Car{}, nil
//}
