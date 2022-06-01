package delivery

import (
	"assignments/layeredcar/model"
	"assignments/layeredcar/service"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type hndleer struct {
	carsevice service.Service
}

func New(serv service.Service) hndleer {
	return hndleer{carsevice: serv}
}

func (h hndleer) Post(w http.ResponseWriter, req *http.Request) {
	y, _ := ioutil.ReadAll(req.Body)
	var data model.Car
	var res model.Car
	var err error
	err = json.Unmarshal(y, &data)
	if err != nil {
		log.Println("Unmarshal error", err)
	}
	res, err = h.carsevice.Create(data)
	if err != nil {
		log.Print("Error adding entries to engine", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	r, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal error", err)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(r))
	if err != nil {
		log.Println("Error writing response")
		return
	}
}
func (h hndleer) Get(w http.ResponseWriter, req *http.Request) {
	path := mux.Vars(req)
	var data model.Car
	k, err := uuid.Parse(path["id"])
	if err != nil {
		log.Println("Error parsing id to UUID")
		w.WriteHeader(http.StatusBadRequest)
	}
	data, err = h.carsevice.GetById(k)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	y, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while marshaling")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(y))

}
func (h hndleer) Put(w http.ResponseWriter, req *http.Request) {
	y, _ := ioutil.ReadAll(req.Body)
	var data model.Car
	var err error
	err = json.Unmarshal(y, &data)
	if err != nil {
		log.Println("Unmarshal error", err)
	}
	path := mux.Vars(req)
	k, err := uuid.Parse(path["id"])
	if err != nil {
		log.Println("Error parsing id to UUID")
		w.WriteHeader(http.StatusBadRequest)
	}
	res, err := h.carsevice.Update(k, data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	r, err := json.Marshal(res)
	if err != nil {
		log.Println("Marshal error", err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(r))
	if err != nil {
		log.Print("Error while writing to response", err)
	}
}
func (h hndleer) Delete(w http.ResponseWriter, req *http.Request) {
	path := mux.Vars(req)
	k, err := uuid.Parse(path["id"])
	if err != nil {
		log.Println("Error parsing id to UUID")
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.carsevice.Delete(k)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusNoContent)

}
func (h hndleer) GetBrand(w http.ResponseWriter, req *http.Request) {
	brand := req.URL.Query().Get("brand")
	engine := req.URL.Query().Get("engine")

	data, err := h.carsevice.GetByBrand(brand, engine)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	y, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while marshaling")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(y))
}
