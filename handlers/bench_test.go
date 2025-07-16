package handlers

import (
	"BasicCRUD/config"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"testing"
)

func BenchmarkCreateCar(b *testing.B) {
	config.ConnectDB()

	app := fiber.New()
	app.Get("/cars/:id", GetCar)

	req, _ := http.NewRequest("GET", "/cars/1", nil)
	req.Header.Set("Content-Type", "application/json")
	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req, 5000)
	}
}
