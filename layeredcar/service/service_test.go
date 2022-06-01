package service

import (
	"assignments/layeredcar/datastore"
	"assignments/layeredcar/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCarServiceCreate(t *testing.T) {
	id1, _ := uuid.NewUUID()
	id2, _ := uuid.NewUUID()
	eInput := model.Engine{Displacement: 100, Noc: 4, Rng: 200}
	eOutput := model.Engine{EngineId: id1, Displacement: 100, Noc: 4, Rng: 200}
	emptycar := model.Car{}
	emptyengine := model.Engine{}
	carinp1 := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eInput}
	carinp2 := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}
	carOutput := model.Car{Id: id2, Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}
	carInvalidBrand := model.Car{Name: "city", Year: 2020, Brand: "Honda", FuelType: "Electric", Engine: eInput}
	carInvalidBrand1 := model.Car{Name: "city", Year: 2020, Brand: "Honda", FuelType: "Electric", Engine: eOutput}
	carInvalidBrandOutput := model.Car{Id: uuid.Nil, Name: "city", Year: 2020, Brand: "Honda", FuelType: "Electric", Engine: eInput}
	carInvalidYear := model.Car{Name: "Tesla Model Y", Year: 2030, Brand: "Tesla", FuelType: "Petrol", Engine: eInput}
	carInvalidYear1 := model.Car{Name: "Tesla Model Y", Year: 2030, Brand: "Tesla", FuelType: "Petrol", Engine: eOutput}
	carInvalidYearOutput := model.Car{Id: uuid.Nil, Name: "Tesla Model Y", Year: 2030, Brand: "Tesla", FuelType: "Petrol", Engine: eInput}
	testcases := []struct {
		desc         string
		carinp1      model.Car
		carinp2      model.Car
		carOutput    model.Car
		engineOutput model.Engine
		expectedErr  error
	}{
		{"Normal entry case", carinp1, carinp2, carOutput, eOutput, nil},
		{"Empty entry case", emptycar, emptycar, emptycar, emptyengine, errors.New("no entry created, Empty data sent")},
		{"Invalid brand", carInvalidBrand, carInvalidBrand1, carInvalidBrandOutput, eOutput, errors.New("no entry created,Invalid Brand")},
		{"Invalid year", carInvalidYear, carInvalidYear1, carInvalidYearOutput, eOutput, errors.New("no entry created,Invalid Year")},
	}
	ctrl := gomock.NewController(t)
	mockCar := datastore.NewMockcar(ctrl)
	mockEngine := datastore.NewMockengine(ctrl)
	for i, tc := range testcases {
		if i == 0 {
			mockEngine.EXPECT().Create(tc.carinp1.Engine).Return(tc.engineOutput, tc.expectedErr)
			mockCar.EXPECT().Create(tc.carinp2).Return(tc.carOutput, tc.expectedErr)
		}
		svc := New(mockCar, mockEngine)
		output, err := svc.Create(tc.carinp1)
		log.Println(err)
		assert.Equal(t, tc.carOutput, output)
		assert.Equal(t, tc.expectedErr, err, "Test failed %v", i)
	}
}
func TestCarServiceUpdate(t *testing.T) {
	id1, _ := uuid.NewUUID()
	eInput := model.Engine{Displacement: 100, Noc: 4, Rng: 200}
	eOutput := model.Engine{EngineId: id1, Displacement: 0, Noc: 0, Rng: 0}
	carinp1 := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eInput}
	carinp2 := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}
	emptycar := model.Car{}
	carInvalidBrand := model.Car{Name: "city", Year: 2020, Brand: "Honda", FuelType: "Electric", Engine: eInput}
	carInvalidBrand1 := model.Car{Name: "city", Year: 2020, Brand: "Honda", FuelType: "Electric", Engine: eOutput}
	carInvalidYear := model.Car{Name: "Tesla Model Y", Year: 2030, Brand: "Tesla", FuelType: "Petrol", Engine: eInput}
	carInvalidYear1 := model.Car{Name: "Tesla Model Y", Year: 2030, Brand: "Tesla", FuelType: "Petrol", Engine: eOutput}
	testcases := []struct {
		desc        string
		id          uuid.UUID
		carinp1     model.Car
		carinp2     model.Car
		carOutput   model.Car
		expectedErr error
	}{
		{"normal update case", id1, carinp1, carinp2, carinp1, nil},
		{"empty entry case", id1, emptycar, emptycar, emptycar, errors.New("no entry updated, Empty data sent")},
		{"Invalid brand", id1, carInvalidBrand, carInvalidBrand1, carInvalidBrand, errors.New("no entry updated,Invalid Brand")},
		{"Invalid year", id1, carInvalidYear, carInvalidYear1, carInvalidYear, errors.New("no entry updated,Invalid Year")},
	}
	ctrl := gomock.NewController(t)
	mockCar := datastore.NewMockcar(ctrl)
	mockEngine := datastore.NewMockengine(ctrl)
	for i, tc := range testcases {
		if i == 0 {
			mockCar.EXPECT().Update(tc.id, tc.carinp1).Return(tc.carinp1, tc.expectedErr)
			mockCar.EXPECT().GetById(tc.id).Return(tc.carinp2, tc.expectedErr)
			mockEngine.EXPECT().Update(tc.carinp2.Engine.EngineId, tc.carinp1.Engine).Return(tc.carinp1.Engine, tc.expectedErr)
		}
		svc := New(mockCar, mockEngine)
		output, err := svc.Update(tc.id, tc.carinp1)
		log.Println(err)
		assert.Equal(t, tc.carOutput, output)
		assert.Equal(t, tc.expectedErr, err, "Test failed %v", i)
	}
}

func TestCarServiceDelete(t *testing.T) {
	id1, _ := uuid.NewUUID()
	eOutput := model.Engine{EngineId: id1, Displacement: 0, Noc: 0, Rng: 0}
	carinp1 := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}
	testcases := []struct {
		desc        string
		id          uuid.UUID
		input       model.Car
		expectedErr error
	}{
		{"Successfully deleted", id1, carinp1, nil},
		{"Empty Id", uuid.Nil, model.Car{}, errors.New("cannot delete,empty id")},
	}
	ctrl := gomock.NewController(t)
	mockCar := datastore.NewMockcar(ctrl)
	mockEngine := datastore.NewMockengine(ctrl)
	for i, tc := range testcases {
		if i == 0 {
			mockCar.EXPECT().GetById(tc.id).Return(tc.input, tc.expectedErr)
			mockCar.EXPECT().Delete(tc.id).Return(tc.expectedErr)
			mockEngine.EXPECT().Delete(tc.input.Engine.EngineId).Return(tc.expectedErr)
		}
		svc := New(mockCar, mockEngine)
		err := svc.Delete(tc.id)
		assert.Equal(t, tc.expectedErr, err, "Test failed %v", i)
	}
}

func TestCarServiceGetById(t *testing.T) {
	id1, _ := uuid.NewUUID()
	eOutput := model.Engine{EngineId: id1, Displacement: 10, Noc: 4, Rng: 200}
	carOutput := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}
	testcases := []struct {
		desc        string
		id          uuid.UUID
		carOutput   model.Car
		engOutput   model.Engine
		expectedErr error
	}{
		{"Successfully fetched", id1, carOutput, eOutput, nil},
		{"Empty Id", uuid.Nil, model.Car{}, model.Engine{}, errors.New("cannot Get,empty id")},
	}
	ctrl := gomock.NewController(t)
	mockCar := datastore.NewMockcar(ctrl)
	mockEngine := datastore.NewMockengine(ctrl)
	for i, tc := range testcases {
		if i == 0 {
			mockCar.EXPECT().GetById(tc.id).Return(tc.carOutput, tc.expectedErr)
			mockEngine.EXPECT().GetById(tc.carOutput.Engine.EngineId).Return(tc.engOutput, tc.expectedErr)
		}
		svc := New(mockCar, mockEngine)
		output, err := svc.GetById(tc.id)
		assert.Equal(t, tc.carOutput, output, "Test failed %v", i)
		assert.Equal(t, tc.expectedErr, err, "Test failed %v", i)
	}
}

func TestCarServiceGetByBrand(t *testing.T) {
	id1, _ := uuid.NewUUID()
	eOutput := model.Engine{EngineId: id1, Displacement: 10, Noc: 4, Rng: 200}
	carOutput := []model.Car{{Id: id1, Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput}}
	eOutput2 := model.Engine{EngineId: id1, Displacement: 0, Noc: 0, Rng: 0}
	carOutput2 := []model.Car{{Id: id1, Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: eOutput2}}
	testcases := []struct {
		desc        string
		brand       string
		engine      string
		carOutput   []model.Car
		engOutput   model.Engine
		expectedErr error
	}{
		{"Successfully fetched", "Ferrari", "included", carOutput, eOutput, nil},
		{"Successfully fetched", "Tesla", "excluded", carOutput2, model.Engine{}, nil},
		{"Empty Brand", "", "", []model.Car{}, model.Engine{}, errors.New("can not get,empty brand")},
	}
	ctrl := gomock.NewController(t)
	mockCar := datastore.NewMockcar(ctrl)
	mockEngine := datastore.NewMockengine(ctrl)
	for i, tc := range testcases {
		if i != 2 {
			mockCar.EXPECT().GetByBrand(tc.brand).Return(tc.carOutput, tc.expectedErr)
			if tc.engine == "included" {
				for _, out := range tc.carOutput {
					mockEngine.EXPECT().GetById(out.Engine.EngineId).Return(tc.engOutput, tc.expectedErr)
				}
			}
		}
		svc := New(mockCar, mockEngine)
		output, err := svc.GetByBrand(tc.brand, tc.engine)
		assert.Equal(t, tc.carOutput, output, "Test failed %v", i)
		assert.Equal(t, tc.expectedErr, err, "Test failed %v", i)
	}
}
