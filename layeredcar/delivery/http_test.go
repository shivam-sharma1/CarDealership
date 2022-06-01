package delivery

import (
	"assignments/layeredcar/model"
	"assignments/layeredcar/service"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Post(t *testing.T) {
	carInput := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: model.Engine{Displacement: 3, Noc: 5, Rng: 100}}
	emptyCar := model.Car{}
	testcases := []struct {
		desc   string
		input  model.Car
		output model.Car
		expErr error
		status int
	}{
		{"Successful entry", carInput, carInput, nil, http.StatusCreated},
		{"empty entry", emptyCar, emptyCar, errors.New("Cannot Post "), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServ := service.NewMockService(ctrl)
	for i, tc := range testcases {
		body, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("Marshall err,test:%v failed", i)
		}
		mockServ.EXPECT().Create(tc.input).Return(tc.output, tc.expErr)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		w := httptest.NewRecorder()
		hndle := New(mockServ)
		hndle.Post(w, req)
		assert.Equal(t, tc.status, w.Result().StatusCode, "Test case failed %v", i)
	}
}

func TestHandler_Put(t *testing.T) {
	carInput := model.Car{Name: "Ferrari Roma", Year: 2019, Brand: "Ferrari", FuelType: "Petrol", Engine: model.Engine{Displacement: 3, Noc: 5, Rng: 100}}
	emptyCar := model.Car{}
	testcases := []struct {
		desc   string
		id     uuid.UUID
		input  model.Car
		output model.Car
		expErr error
		status int
	}{
		{"normal entry case", uuid.New(), carInput, carInput, nil, http.StatusOK},
		{"empty entry case", uuid.Nil, emptyCar, emptyCar, errors.New("Cannot Put "), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServ := service.NewMockService(ctrl)
	for i, tc := range testcases {
		body, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("Marshall err,test:%v failed", i)
		}
		mockServ.EXPECT().Update(tc.id, tc.input).Return(tc.output, tc.expErr)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(body))
		req = mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()
		hndle := New(mockServ)
		hndle.Put(w, req)
		assert.Equal(t, tc.status, w.Result().StatusCode, "Test case failed %v", i)
	}
}
func TestHandler_GetId(t *testing.T) {
	output1 := model.Car{uuid.New(), "Ferrari Roma", 2019, "Ferrari", "Petrol", model.Engine{uuid.New(), 3, 5, 100}}
	testcases := []struct {
		desc   string
		id     uuid.UUID
		output model.Car
		expErr error
		status int
	}{
		{"passed id", uuid.New(), output1, nil, http.StatusOK},
		{"empty id", uuid.Nil, model.Car{}, errors.New("can not get"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServ := service.NewMockService(ctrl)
	for i, tc := range testcases {
		mockServ.EXPECT().GetById(tc.id).Return(tc.output, tc.expErr)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()
		hndle := New(mockServ)
		hndle.Get(w, req)
		assert.Equal(t, tc.status, w.Result().StatusCode, "Test case failed %v", i)
	}
}

func TestHandler_GetBrand(t *testing.T) {
	output1 := []model.Car{{uuid.New(), "Ferrari Roma", 2019, "Ferrari", "Petrol", model.Engine{uuid.New(), 3, 5, 100}}}
	testcases := []struct {
		desc   string
		brand  string
		engine string
		output []model.Car
		expErr error
		status int
	}{
		{"passed brand", "Ferrari", "included", output1, nil, http.StatusOK},
		{"empty brand", "", "", []model.Car{}, errors.New("can not get"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServ := service.NewMockService(ctrl)
	for i, tc := range testcases {
		mockServ.EXPECT().GetByBrand(tc.brand, tc.engine).Return(tc.output, tc.expErr)
		req := httptest.NewRequest(http.MethodGet, "/?brand="+tc.brand+"&engine="+tc.engine, nil)
		w := httptest.NewRecorder()
		hndle := New(mockServ)
		hndle.GetBrand(w, req)
		assert.Equal(t, tc.status, w.Result().StatusCode, "Test case failed %v", i)
	}
}
func TestHandler_Delete(t *testing.T) {
	testcases := []struct {
		desc   string
		id     uuid.UUID
		expErr error
		status int
	}{
		{"Successful deletion", uuid.New(), nil, http.StatusNoContent},
		{"empty entry ", uuid.Nil, errors.New("can not delete"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServ := service.NewMockService(ctrl)

	for i, tc := range testcases {
		mockServ.EXPECT().Delete(tc.id).Return(tc.expErr)
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		hndle := New(mockServ)
		hndle.Delete(w, req)

		assert.Equal(t, tc.status, w.Result().StatusCode, "Test case failed %v", i)
	}
}

//package delivery
//
//import (
//	"assignments/layeredcar/models"
//	"github.com/google/uuid"
//	"testing"
//)
//
//func TestPost(t *testing.T) {
//	tc := []struct {
//		desc string
//		inp  models.Car
//		out  models.Car
//		err  error
//	}{
//		{"Successful input", models.Car{uuid.New(), "Mercedes Benz GLA", 2018, "Mercedes", models.Engine{uuid.New(), uuid.New(), 400, 4, nil}, "Petrol"}, models.Car{uuid.New(), "Mercedes Benz GLA", 2018, "Mercedes", models.Engine{uuid.New(), uuid.New(), 400, 4, 0}, "Petrol"}, nil},
//		{"old car model"},
//	}
//}
