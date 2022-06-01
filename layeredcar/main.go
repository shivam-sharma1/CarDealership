package main

import (
	"assignments/layeredcar/datastore/car"
	"assignments/layeredcar/datastore/engine"
	"assignments/layeredcar/delivery"
	"assignments/layeredcar/service"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ConnectToMySql() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Root@123"
	dbName := "cardealership"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		fmt.Println("cannot connect")
		return nil, err
	}
	fmt.Println("connected")
	return db, nil
}

func main() {
	db, err := ConnectToMySql()
	if err != nil {
		log.Println("Database not connected!", err)
		return
	}
	carStore := car.New(db)
	engineStore := engine.New(db)
	service := service.New(carStore, engineStore)
	handl := delivery.New(service)

	r := mux.NewRouter()
	r.HandleFunc("/", handl.Post)
	r.HandleFunc("/{id}", handl.Put)
	r.HandleFunc("/", handl.GetBrand)
	r.HandleFunc("/{id}", handl.Get)
	r.HandleFunc("/{id}", handl.Delete)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
