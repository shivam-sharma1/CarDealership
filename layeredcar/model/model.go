package model

import "github.com/google/uuid"

type Engine struct {
	EngineId     uuid.UUID `json:"engineId"`
	Displacement int       `json:"displacement"`
	Noc          int       `json:"noc"`
	Rng          int       `json:"rng"`
}
type Car struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Year     int       `json:"year"`
	Brand    string    `json:"brand"`
	FuelType string    `json:"fuelType"`
	Engine   Engine    `json:"engineId"`
}

//uuid.UUID

//package models
//
//import (
//	"github.com/google/uuid"
//)
//
//type Car struct {
//	Id       uuid.UUID `json:"id"json:"id"`
//	Name     string    `json:"name"`
//	Year     int       `json:"year,omitempty"`
//	Brand    string    `json:"brand"`
//	Engine   Engine    `json:"engine,omitempty"`
//	FuelType string    `json:"fueltype,omitempty"`
//}
