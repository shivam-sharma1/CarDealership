package car

import (
	"assignments/layeredcar/model"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCarstorer_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	testcases := []struct {
		desc  string
		input model.Car
		err   error
	}{
		{"Successfully created", model.Car{uuid.Nil, "Tesla Model X", 2019, "Tesla", "Petrol", model.Engine{EngineId: uuid.New()}}, nil},
		{"Error case", model.Car{}, errors.New("no entry created")},
	}

	for i, tc := range testcases {
		mock.ExpectExec(InsertQuery).WithArgs(sqlmock.AnyArg(), tc.input.Name, tc.input.Year, tc.input.Brand, tc.input.FuelType, tc.input.Engine.EngineId).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		d := New(db)
		_, err := d.Create(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

func TestCarstorer_GetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	id := uuid.New()
	testcases := []struct {
		desc  string
		input uuid.UUID
		err   error
	}{
		{"Successfully fetched", id, nil},
		{"Error case", uuid.Nil, errors.New("no entry fetched")},
	}

	rows := mock.NewRows([]string{"id", "name", "year", "brand", "fueltype", "engineid"}).AddRow(id.String(), "Tesla Model X", 2020, "Tesla", "petrol", uuid.New())
	for i, tc := range testcases {
		mock.ExpectQuery(SelectByIdQuery).WithArgs(tc.input).
			WillReturnRows(rows).WillReturnError(tc.err)
		d := New(db)
		_, err := d.GetById(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}
func TestCarstorer_GetByBrand(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	testcases := []struct {
		desc  string
		input string
		err   error
	}{
		{"Successfully fetched", "Tesla", nil},
		{"Error case", "", errors.New("no entry fetched")},
	}

	rows := mock.NewRows([]string{"id", "name", "year", "brand", "fueltype", "engineid"}).AddRow(uuid.NewString(), "Tesla Model X", 2020, "Tesla", "petrol", uuid.NewString())
	for i, tc := range testcases {
		mock.ExpectQuery(SelectByBrandQuery).WithArgs(tc.input).
			WillReturnRows(rows).WillReturnError(tc.err)
		d := New(db)
		_, err := d.GetByBrand(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}
func TestCarstorer_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	id := uuid.New()
	testcases := []struct {
		desc  string
		input model.Car
		err   error
	}{
		{"Successfully updated", model.Car{id, "Tesla Model X", 2020, "Tesla", "Petrol", model.Engine{EngineId: uuid.New()}}, nil},
		{"Error case", model.Car{}, errors.New("no entry updated")},
	}

	for i, tc := range testcases {
		mock.ExpectExec(UpdateQuery).WithArgs(tc.input.Name, tc.input.Year, tc.input.Brand, tc.input.FuelType, tc.input.Id).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		d := New(db)
		_, err := d.Update(tc.input.Id, tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

func TestCarstorer_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	id := uuid.New()
	testcases := []struct {
		desc  string
		input uuid.UUID
		err   error
	}{
		{"Successfully deleted", id, nil},
		{"Error case", uuid.Nil, errors.New("no entry deleted")},
	}

	for i, tc := range testcases {
		mock.ExpectExec(DeleteQuery).WithArgs(tc.input).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		d := New(db)
		err := d.Delete(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

//First try
//package car
//
//import (
//	"assignments/layeredcar/model"
//	"context"
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/google/uuid"
//	"log"
//	"reflect"
//	"testing"
//)
//
//func NewMock() (*sql.DB, sqlmock.Sqlmock) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		log.Fatalf("an error '%s' was not expected when opening the database connection", err)
//	}
//
//	return db, mock
//}
//func TestCarStorer_Post(t *testing.T) {
//	tempid := uuid.New()
//	tempengineid := uuid.New()
//	car := models.Car{Id: tempid, Name: "BMW X5", Year: 2019, Brand: "BMW", FuelType: "Petrol", Engine: models.Engine{Engineid: tempengineid, Car_id: tempid, Displacement: 2000, Noc: 4, Rnge: 170}}
//	err := errors.New("query error")
//	tc := []struct {
//		desc  string
//		input models.Car
//		err   error
//		out   models.Car
//	}{
//		{desc: "Successfully entered values", input: models.Car{Id: tempid, Name: "BMW X5", Year: 2019, Brand: "BMW", Engine: models.Engine{Engineid: tempengineid, Car_id: tempid, Displacement: 2000, Noc: 4, Rnge: 170}, FuelType: "Petrol"}, out: models.Car{Id: tempid, Name: "BMW X5", Year: 2019, Brand: "BMW", Engine: models.Engine{Engineid: tempengineid, Car_id: tempid, Displacement: 2000, Noc: 4, Rnge: 170}, FuelType: "Petrol"}},
//		{"Empty values", models.Car{}, errors.New("no entry created"), models.Car{}},
//	}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	a := New(db)
//	if err != nil {
//		t.Error(err)
//	}
//	defer db.Close()
//
//	mock.ExpectExec("insert into cars values(?,?,?,?,?,?)").
//		WithArgs(car.Id, car.Engine.Engineid, car.Name, car.Year, car.Brand, car.FuelType).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	mock.ExpectExec("insert into cars values(?,?,?,?,?,?)").
//		WithArgs(car.Id, car.Engine.Engineid, car.Name, car.Year, car.Brand, car.FuelType).
//		WillReturnError(err)
//
//	for i, tc := range tc {
//		res, err := a.CarStorer_Post(context.TODO(), &car)
//
//		if res != tc.out {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, res, tc.out)
//		}
//
//		if !reflect.DeepEqual(err, tc.err) {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, err, tc.err)
//		}
//	}
//}
//
//func TestCarStorer_Put(t *testing.T) {
//	tempcarid := uuid.New()
//	tempengineid := uuid.New()
//	var (
//		car = models.Car{Id: tempcarid, Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{Engineid: tempengineid, Car_id: tempcarid, Displacement: 250, Noc: 4, Rnge: 150}, FuelType: "Electric"}
//		//car2 = models.Car{Id: id, Name: "Tesla Model S", Year: 2020, Brand: "Tesla", FuelType: "electric", Engine: models.Engine{}}
//	)
//	tc := []struct {
//		desc string
//		Id   uuid.UUID
//		car  models.Car
//		err  error
//		out  models.Car
//	}{
//		{"Update values correctly", tempcarid, models.Car{Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{tempengineid, tempcarid, 250, 4, 150}, FuelType: "Electric"}, nil, models.Car{tempcarid, "Tesla Model X", 2021, "Tesla", models.Engine{tempengineid, tempcarid, 250, 4, 150}, "Electric"}},
//		//{"No change in values", tempcarid, models.Car{Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{tempengineid, tempcarid, 250, 4, 150}, FuelType: "Electric"}, errors.New("no new values"), models.Car{tempcarid, "Tesla Model X", 2021, "Tesla", models.Engine{tempengineid, tempcarid, 250, 4, 150}, "Electric"}},
//		//{"car does not exist", tempcarid, models.Car{Name: "Tata Nano", Year: 2022, Brand: "Tata", Engine: models.Engine{uuid.New(), uuid.New(), 400, 4, 150}, FuelType: "Electric"}, errors.New("car does not exist"), models.Car{}},
//	}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	if err != nil {
//		t.Error(err)
//	}
//	a := New(db)
//	defer db.Close()
//	mock.ExpectExec("update cars set name=?,year=?,brand=?,fueltype=? where id=?").
//		WithArgs(car.Name, car.Year, car.Brand, car.FuelType, car.Id).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	//mock.ExpectExec("update cars set name=?,year=?,brand=?,fuel_type=? WHERE id=?").
//	//	WithArgs(car.Name, car.Year, car.Brand, car.FuelType, car.Id).
//	//	WillReturnError(err)
//
//	for i, tc := range tc {
//		newCar, err := a.CarStorer_Put(context.TODO(), tc.Id.String(), &car)
//		if !reflect.DeepEqual(newCar, tc.out) {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, newCar, tc.out)
//		}
//		if !reflect.DeepEqual(err, tc.err) {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, err, tc.err)
//		}
//	}
//
//}
//func TestCarStorer_GetID(t *testing.T) {
//	var (
//		tempcarid    = uuid.New()
//		tempengineid = uuid.New()
//		//id1, _ = uuid.Parse("fa872ede-fc16-4d98-9a0a-d70fae3342dc")
//		car1 = models.Car{Id: tempcarid, Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{
//			Engineid: tempengineid, Car_id: tempcarid, Displacement: 120, Noc: 4, Rnge: 150}, FuelType: "Electric"}
//		err = errors.New("element by id failed")
//	)
//	tc := []struct {
//		desc string
//		id   uuid.UUID
//		err  error
//		out  models.Car
//	}{
//		{desc: "id found ", id: tempcarid, out: models.Car{Id: uuid.New(), Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{Engineid: tempengineid, Car_id: tempcarid, Displacement: 120, Noc: 4, Rnge: 150}, FuelType: "Electric"}},
//		{"id does not exist", tempcarid, errors.New("entered id does not exist"), models.Car{}},
//	}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	if err != nil {
//		log.Println(err)
//	}
//	defer db.Close()
//	a := New(db)
//	rows := sqlmock.NewRows([]string{"id", "engineId", "name", "year", "brand", "fuelType"}).
//		AddRow(car1.Id.String(), car1.Engine.Engineid, car1.Name, car1.Year, car1.Brand, car1.FuelType)
//	mock.ExpectQuery("select * from cars where id=?").WithArgs(tempcarid).WillReturnRows(rows)
//	for i, tc := range tc {
//		resp, err := a.CarStorer_GetID(context.TODO(), tc.id.String())
//		fmt.Println(resp)
//		if resp != tc.out {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, resp, tc.out)
//		}
//		if err != tc.err {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, err, tc.err)
//		}
//	}
//}
//
//func TestCarStorer_GetBrand(t *testing.T) {
//	var (
//		id1, _ = uuid.Parse("fa872ede-fc16-4d98-9a0a-d70fae3342dc")
//		id2, _ = uuid.Parse("0898ebe6-c896-427c-a6e3-f0ed2eb43507")
//		err    = errors.New("error")
//		car1   = models.Car{Id: id1, Name: "Tesla Model X", Year: 2021, Brand: "Tesla",
//			FuelType: "electric", Engine: models.Engine{}}
//
//		car2 = models.Car{Id: id2, Name: "tesla Model Y", Year: 2022, Brand: "Tesla",
//			FuelType: "electric", Engine: models.Engine{}}
//	)
//	tc := []struct {
//		desc  string
//		brand string
//		err   error
//		out   models.Car
//	}{
//		{desc: "Brand found", brand: "Tesla", out: models.Car{Id: uuid.New(), Name: "Tesla Model X", Year: 2021, Brand: "Tesla", Engine: models.Engine{Engineid: uuid.New(), Car_id: uuid.New(), Displacement: 300, Noc: 4, Rnge: 150}, FuelType: "Electric"}},
//		{"brand not found", "Suzuki", errors.New("given brand does not exist"), models.Car{}},
//	}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	if err != nil {
//		t.Error(err)
//	}
//	a := New(db)
//	defer db.Close()
//	rows := sqlmock.NewRows([]string{"id", "engineId", "name", "year", "brand", "fuelType"}).
//		AddRow(id1.String(), id1.String(), car1.Name, car1.Year, car1.Brand, car1.FuelType).
//		AddRow(id2.String(), id2.String(), car2.Name, car2.Year, car2.Brand, car2.FuelType)
//	mock.ExpectQuery("select * from cars where brand=?").WithArgs("Suzuki").WillReturnRows(rows)
//	mock.ExpectQuery("select * from cars where brand=?").WithArgs("").WillReturnError(err)
//
//	for i, tc := range tc {
//		car, err := a.GetCarsByBrand(context.TODO(), tc.brand)
//		if !reflect.DeepEqual(err, tc.err) {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, err, tc.err)
//		}
//
//		if !reflect.DeepEqual(car, tc.out) {
//			t.Errorf("\n[TEST %v] Failed \n got %v\nGot \n Expected %v", i, car, tc.out)
//		}
//	}
//}
//func TestCarStorer_Delete(t *testing.T) {
//	id1, _ := uuid.Parse("fa872ede-fc16-4d98-9a0a-d70fae3342dc")
//	err := errors.New("delete failed")
//	tc := []struct {
//		desc string
//		id   uuid.UUID
//		err  error
//	}{
//		{"deleted successfully", uuid.New(), nil},
//		{"value not found", uuid.New(), errors.New("value not found")},
//	}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	if err != nil {
//		t.Error(err)
//	}
//	a := New(db)
//	defer db.Close()
//	mock.ExpectExec("DELETE FROM cars WHERE ID=?").WithArgs(id1.String()).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectExec("DELETE FROM cars WHERE ID=?").WithArgs(uuid.Nil).
//		WillReturnError(err)
//	for i, tc := range tc {
//		_, err := a.DeleteCar(context.TODO(), tc.id.String())
//
//		if err != tc.err {
//			t.Errorf("\n[TEST %v] Failed \nDesc %v\nGot %v\n Expected %v", i, tc.desc, err, tc.err)
//		}
//	}
//}
