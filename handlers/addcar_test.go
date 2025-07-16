package handlers

import (
	"BasicCRUD/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestCreateCar(t *testing.T) {
	config.ConnectDB()

	app := fiber.New()
	app.Post("/cars", CreateCar)
	body := `
{
	"name": "Corolla",
	"model": "xt",
	"brand": "Totoya",
	"year" : 2022,
	"price": 302034
}
`
	req, _ := http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Request Failed : %v", err)
	}
	assert.Equalf(t, fiber.StatusCreated, resp.StatusCode, "Response should be OK")
}
