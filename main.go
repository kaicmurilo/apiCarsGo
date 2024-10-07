package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// Car struct representa um carro com os campos Name e Price
type Car struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Declarando uma variável global para armazenar uma lista de carros
var cars []Car

// busca todos os carros do banco de dados e retorna em formato json
func getCars(c echo.Context) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, price FROM cars")
	if err != nil {
		return err
	}
	defer rows.Close()

	cars = nil // Limpa o slice `cars`
	for rows.Next() {
		var car Car
		err = rows.Scan(&car.Name, &car.Price)
		if err != nil {
			return err
		}
		cars = append(cars, car)
	}

	return c.JSON(200, cars)
}

// cria um novo carro no banco de dados
func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	saveCar(*car)
	return c.JSON(201, *car)
}

// insere um carro no banco de dados SQLite
func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars(name, price) values($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}

	return nil
}

// inicializa o servidor e define as rotas para a API
func main() {
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}
