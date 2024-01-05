package controller

import (
	"fmt"
	"test/models"
	"time"

	"github.com/google/uuid"
)

func getCarInfo() models.Car{
	var (
		model, brand, idStr string
		cmd, year int
	)

	a:
		fmt.Print(`Enter command:
				1 - Create
				2 - Update
		`)
		fmt.Scan(&cmd)

		if cmd == 2 {
			fmt.Print("Enter id: ")
			fmt.Scan(&idStr)
			fmt.Println()

			fmt.Print("Enter model and brand: ")
			fmt.Scan(&model, &brand)
			fmt.Println()

			fmt.Print("Enter year: ")
			fmt.Scan(&year)
			
		}else if cmd == 1 {
			fmt.Print("Enter model and brand: ")
			fmt.Scan(&model, &brand)
			fmt.Println()

			fmt.Print("Enter year: ")
			fmt.Scan(&year)

		}else{
			fmt.Println("Not found!")
			goto a
		}
		if idStr != ""{
			return models.Car{
				ID:		uuid.MustParse(idStr),
				Model:	model,
				Brand:	brand,
				Year:	year,
			}
		}
		return models.Car{
			Model: model,
			Brand: brand,
			Year: year,
		}

}

func (c Controller) CreateCar()  {
	car := getCarInfo()

	if car.Year <= 0 || car.Year > time.Now().Year()+1{
		fmt.Println("Year input is not correct!")
		return
	}

	id, err := c.Store.CarStorage.Insert(car)
	if err != nil{
		fmt.Println("Error while inserting data inside controller!", err)
		return
	}
	fmt.Println("id: ", id)
}

func (c Controller) GetCarByID()  {
	idStr := ""
	fmt.Print("Enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Print("Id is not uuid error: ", err.Error())
		return
	}

	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil{
		fmt.Println("Error while getting car by id! :",err.Error())
		return
	}
	fmt.Println("Your car is:", car)
}

func (c Controller) GetCarList(){
	cars, err := c.Store.CarStorage.GetList()
	if err != nil{
		fmt.Println("Error while getting list: ", err.Error())
		return
	}
	fmt.Println(cars)
}

func (c Controller) UpdateCar (){
	car := getCarInfo()

	if !checkCarInfo(car){
		return
	}

	err := c.Store.CarStorage.Update(car)
	if err != nil{
		fmt.Println("Error while updating car: ", err.Error())
		return
	}
	if car.ID.String() != ""{
		fmt.Println("Successfully updated!")
	}else{
		fmt.Println("Successfullu created!")
	}
}

func checkCarInfo(car models.Car) bool{
	if car.Year <= 0 || car.Year > time.Now().Year()+1{
		fmt.Println("Year input is not correct!")
		return false
	}
	return true
}