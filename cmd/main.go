package main

import (
	"log"
	"test/config"
	"test/controller"
	"test/storage/postgres"
)

func main()  {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil{
		log.Fatalln("Error while connecting to db err: ", err.Error())
		return
	}
	defer store.DB.Close()

	 con := controller.New(store)

	// CARS

	con.CreateCar()
	// con.GetCarList()
	// con.UpdateCar()
	//con.GetCarByID()

	// DRIVERS

	// con.CreateDriver()
	// con.GetDriverByID()
	// con.GetDriverList()
	// con.UpdateDriver()
}