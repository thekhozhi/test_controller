package postgres

import (
	"database/sql"
	"test/models"

	"github.com/google/uuid"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB)driverRepo{
	return driverRepo{
		DB:	db,
	}
}

func (d driverRepo) Insert ( driver models.Driver) (string, error) {
	id := uuid.New()
	_, err := d.DB.Exec(`INSERT INTO drivers values ($1, $2, $3)`,id, driver.FullName, driver.Phone)
	if err != nil{
		return 	"", err
	}
	return id.String(),nil
}

func (d driverRepo) GetById(id uuid.UUID)(models.Driver, error){
	driver := models.Driver{}
	err := d.DB.QueryRow(`SELECT id, fullname, phone FROM drivers where id = $1`,id).Scan(&driver.ID, &driver.FullName, &driver.Phone)
	if err != nil{
		return models.Driver{}, err
	}
	return driver, nil
}

func (d driverRepo) GetList()([]models.Driver, error){
	rows, err := d.DB.Query(`SELECT * FROM drivers`)
	if err != nil{
		return nil, err
	}
	drivers := []models.Driver{}
	for rows.Next(){
		driver := models.Driver{}
		err := rows.Scan(&driver.ID, &driver.FullName, &driver.Phone)
		if err != nil{
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (d driverRepo) Update( driver models.Driver) error{
	_, err := d.DB.Exec(`UPDATE drivers set fullname = $1, phone = $2`,driver.FullName, driver.Phone)
	if err != nil{
		return err
	}
	return nil
}