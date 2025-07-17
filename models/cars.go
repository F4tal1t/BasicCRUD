package models

import (
	"BasicCRUD/config"
	"errors"
	"fmt"
)

type Car struct {
	ID    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

// Insert adds the car to DB and updates the ID
func (c *Car) Insert() error {
	// OLD INSERT OPERATION USING RAW SQL
	// query := `
	//     INSERT INTO cars (name, model, brand, year, price)
	//     VALUES ($1, $2, $3, $4, $5)
	//     RETURNING id
	// `
	// err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.ID)
	// if err != nil {
	//     return fmt.Errorf("insert error: %v", err)
	// }

	// NEW GORM INSERT OPERATION
	result := config.DB.Create(c)
	if result.Error != nil {
		return fmt.Errorf("failed to insert car: %v", result.Error)
	}
	return nil
}

// GetByID fetches car from DB using ID
func (c *Car) GetByID(id int) error {
	// OLD GET OPERATION USING RAW SQL
	// query := `
	//     SELECT id, name, model, brand, year, price
	//     FROM cars
	//     WHERE id = $1
	// `
	// row := config.DB.QueryRow(query, id)
	// err := row.Scan(&c.ID, &c.Name, &c.Model, &c.Brand, &c.Year, &c.Price)
	// if err != nil {
	//     return fmt.Errorf("car not found: %v", err)
	// }

	// NEW GORM GET OPERATION
	result := config.DB.First(c, id)
	if result.Error != nil {
		return fmt.Errorf("car not found: %v", result.Error)
	}
	return nil
}

// Update modifies car details for given ID
func (c *Car) Update(id int) error {
	// OLD UPDATE OPERATION USING RAW SQL
	// query := `
	//     UPDATE cars
	//     SET name = $1, model = $2, brand = $3, year = $4, price = $5
	//     WHERE id = $6
	// `
	// result, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, id)
	// if err != nil {
	//     return fmt.Errorf("update error: %v", err)
	// }
	// rowsAffected, _ := result.RowsAffected()
	// if rowsAffected == 0 {
	//     return errors.New("no car found to update")
	// }
	// c.ID = id

	// NEW GORM UPDATE OPERATION
	c.ID = id
	result := config.DB.Save(c)
	if result.Error != nil {
		return fmt.Errorf("failed to update car: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("no car found to update")
	}
	return nil
}

// DeleteCar removes car from DB by ID
func DeleteCar(id int) error {
	// OLD DELETE OPERATION USING RAW SQL
	// query := `DELETE FROM cars WHERE id = $1`
	// result, err := config.DB.Exec(query, id)
	// if err != nil {
	//     return fmt.Errorf("delete error: %v", err)
	// }
	// rowsAffected, _ := result.RowsAffected()
	// if rowsAffected == 0 {
	//     return errors.New("no car found to delete")
	// }

	// NEW GORM DELETE OPERATION
	result := config.DB.Delete(&Car{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete car: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("no car found to delete")
	}
	return nil
}
