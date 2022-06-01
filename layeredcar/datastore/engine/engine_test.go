package engine

import (
	"assignments/layeredcar/model"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestEnginestorer_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	testcases := []struct {
		desc  string
		input model.Engine
		err   error
	}{
		{"Successfully created", model.Engine{uuid.Nil, 1, 4, 100}, nil},
		{"Error case", model.Engine{}, errors.New("no entry created")},
	}
	for i, tc := range testcases {
		mock.ExpectExec(InsertQuery).WithArgs(sqlmock.AnyArg(), tc.input.Displacement, tc.input.Noc, tc.input.Rng).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		d := New(db)
		_, err := d.Create(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

func TestEnginestorer_GetById(t *testing.T) {
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

	rows := mock.NewRows([]string{"id", "displacement", "noc", "rng"}).AddRow(id, 1, 4, 100)
	for i, tc := range testcases {
		mock.ExpectQuery(SelectByIdQuery).WithArgs(tc.input).
			WillReturnRows(rows).WillReturnError(tc.err)
		d := New(db)
		_, err := d.GetById(tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

func TestEnginestorer_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println("error establishing mock connection", err)
		return
	}
	id := uuid.New()
	testcases := []struct {
		desc  string
		input model.Engine
		err   error
	}{
		{"Successfully updated", model.Engine{id, 1, 4, 100}, nil},
		{"Error case", model.Engine{}, errors.New("no entry updated")},
	}

	for i, tc := range testcases {
		mock.ExpectExec(UpdateQuery).WithArgs(tc.input.Displacement, tc.input.Noc, tc.input.Rng, tc.input.EngineId).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		d := New(db)
		_, err := d.Update(tc.input.EngineId, tc.input)
		assert.Equal(t, tc.err, err, "Test failed %v", i)
	}
}

func TestEnginestorer_Delete(t *testing.T) {
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

//package engine
//
//import (
//	"errors"
//	"github.com/google/uuid"
//	"testing"
//)
//
//func TestEngine_Post(t *testing.T) {
//	tc := []struct {
//		desc  string
//		input models.Engine
//		err   error
//		out   models.Engine
//	}{
//		{"Successfully entered values", models.Engine{uuid.New(), uuid.New(), 2000, 4, 170}, nil, models.Engine{uuid.New(), uuid.New(), 2000, 4, 170}},
//		{"Empty values", models.Engine{}, errors.New("no entry created"), models.Engine{}},
//	}
//}
//
//func TestEngine_Put(t *testing.T) {
//	tc := []struct {
//		desc  string
//		carId uuid.UUID
//		car   models.Engine
//		err   error
//		out   models.Engine
//	}{
//		{"Update values correctly", uuid.New(), models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}, nil, models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}},
//		{"No change in values", uuid.New(), models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}, errors.New("no new values"), models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}},
//		{"engine does not exist", uuid.New(), models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}, errors.New("engine does not exist"), models.Engine{}},
//	}
//}
//func TestEngine_GetID(t *testing.T) {
//	tc := []struct {
//		desc string
//		id   uuid.UUID
//		err  error
//		out  models.Engine
//	}{
//		{"id found ", uuid.New(), nil, models.Engine{uuid.New(), uuid.New(), 2000, 4, 150}},
//		{"id does not exist", uuid.New(), errors.New("entered id does not exist"), models.Engine{}},
//	}
//}
//
////get by brand is not needed in engine
////func TestEngine_GetBrand(t *testing.T){
////	tc := []struct {
////		desc  string
////		brand string
////		err   error
////	}{
////		{"Brand found", "Tesla", nil},
////		{"brand not found", "Suzuki", errors.New("given brand does not exist")},
////	}
////}
//func TestEngine_Delete(t *testing.T) {
//	tc := []struct {
//		desc  string
//		carId uuid.UUID
//		err   error
//	}{
//		{"deleted successfully", uuid.New(), nil},
//		{"value not found", uuid.New(), errors.New("value not found")},
//	}
//}
