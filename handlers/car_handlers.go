package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"sync"

	"BasicCRUD/models"
)

var mu sync.Mutex

// @Summary      Create a new car
// @Description  Add a new car to the inventory
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        car  body  models.Car  true  "Car details"
// @Success      200  {object}  models.Car
// @Router       /cars [post]
func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Incorrect Input Body",
			"details": err.Error(),
		})
	}

	err := car.Insert()
	if err != nil {
		return err
	}

	fmt.Println("Car saved to the inventory with the id:", car.ID)
	return c.Status(fiber.StatusCreated).JSON(car)
}

// @Summary      Get a car by ID
// @Description  Retrieve a car from the inventory by its ID
// @Tags         Cars
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  models.Car
// @Router       /cars/{id} [get]
func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	if err := car.GetByID(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Car not found",
			"id":    car.ID,
		})
	}

	return c.Status(fiber.StatusOK).JSON(car)
}

// @Summary      Update a car
// @Description  Update an existing car in the inventory
// @Tags         Cars
// @Accept       json
// @Produce      json
// @Param        id   path      int         true  "Car ID"
// @Param        car  body      models.Car  true  "Updated car details"
// @Success      200  {object}  models.Car
// @Router       /cars/{id} [put]
func UpdateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	updatedCar := &models.Car{}
	if err := c.BodyParser(updatedCar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Incorrect Input Body",
			"details": err.Error(),
		})
	}

	if err := updatedCar.Update(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update car",
			"details": err.Error(),
		})
	}

	fmt.Println("Car updated successfully with ID:", id)
	return c.Status(fiber.StatusOK).JSON(updatedCar)
}

// @Summary      Delete a car
// @Description  Remove a car from the inventory by its ID
// @Tags         Cars
// @Param        id   path      int  true  "Car ID"
// @Success      204  "No Content"
// @Router       /cars/{id} [delete]
func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	car := &models.Car{}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	car.ID = id
	models.DeleteCar(id)

	fmt.Println("Car deleted successfully")
	return c.SendStatus(fiber.StatusNoContent)
}

func extractIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("missing ID")
	}
	return strconv.Atoi(parts[2])
}
