package controller

import (
	"fmt"
	"test/models"

	"github.com/google/uuid"
)

func getDriverInfo() models.Driver{
	var(
		fullName, phone string
	)
	fmt.Print("Enter driver's full name: ")
	fmt.Scan(&fullName)

	fmt.Println("Enter driver's phone number:")
	fmt.Scan(&phone)

	return models.Driver{
		FullName: fullName,
		Phone: phone,
	}
}

func checkPhoneNumber(phone string) bool {
	for _, r := range phone {
		if r > '0' || r < '9' || r != '+'{
			return false
		}
	}
	return true
}

func (c Controller) CreateDriver(){
	driver := getDriverInfo()

	if !checkPhoneNumber(driver.Phone){
		fmt.Println("The phone number format is not correct!")
		return
	}
	id, err := c.Store.DriverStorage.Insert(driver)
	if err != nil{
		fmt.Println("Error while Inserting driver: ", err.Error())
		return
	}
	fmt.Println("Your new driver's id is:", id)
}

func (c Controller) GetDriverByID(){
	idStr := ""
	fmt.Print("Enter driver's id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil{
		fmt.Println("Id is not uuid: ", err.Error())
		return 
	}

	driver, err := c.Store.DriverStorage.GetById(id)
	if err != nil{
		fmt.Println("Error while selecting by id: ", err.Error())
		return 
	}
	fmt.Println("The driver is: ", driver)
}

func (c Controller)GetDriverList(){
	drivers, err := c.Store.DriverStorage.GetList()
	if err != nil{
		fmt.Println("Error while selecting all data from drivers:", err.Error())
		return
	}
	fmt.Println(drivers)
}

func (c Controller) UpdateDriver() {
	driver := getDriverInfo()

	err := c.Store.DriverStorage.Update(driver)
	if err != nil{
		fmt.Println("Error while updating driver: ", err.Error())
	}
	if driver.ID.String() != ""{
		fmt.Println("Successfully Updated!")
	} else {
		fmt.Println("Successfully Created!")
	}
}
