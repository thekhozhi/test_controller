package postgres

import (
	"database/sql"
	"fmt"
	"test/models"

	"github.com/google/uuid"
)

type carRepo struct {
	DB *sql.DB
}

func NewCarRepo(db *sql.DB) carRepo{
	return carRepo{
		DB: db,
	}
}

func (c carRepo) Insert(car models.Car)  (string, error) {
	id := uuid.New()

	_, err := c.DB.Exec(`INSERT INTO cars values ($1, $2, $3, $4)`, id, car.Model, car.Brand, car.Year)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (c carRepo) GetByID(id uuid.UUID) (models.Car, error){
car := models.Car{}

	err := c.DB.QueryRow(`SELECT from table where id = $1`, id).Scan(
		&car.ID,
		&car.Model,
		&car.Brand,
		&car.Year,
	)
	if err != nil{
		fmt.Println("Error while selecting car by id!", err.Error())
		return models.Car{}, err
	}
	return car, nil
}

func (c carRepo) GetList() ([]models.Car, error) {
	cars := []models.Car{}

	rows, err := c.DB.Query(`SELECT * FROM cars`)
	if err != nil{
		 return nil, err
	}
	 for rows.Next(){
		car := models.Car{}
		err := rows.Scan(&car.ID, &car.Model, &car.Brand, &car.Year)
		if err != nil{
			return nil, err
		}
		cars = append(cars, car)
	 }
	 return cars, nil
}

func (c carRepo) Update(car models.Car) error {
	_, err := c.DB.Exec(`UPDATE cars set model = $1, brand = $2, year = $3 where id = $4`, car.Model, car.Brand, car.Year, car.ID)
	if err != nil{
		return err
	}
	return nil
}