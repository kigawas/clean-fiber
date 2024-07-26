package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/clean-fiber/api"
	"github.com/kigawas/clean-fiber/app"
	"github.com/kigawas/clean-fiber/models/schemas"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	config := app.Config{
		DatabaseURL: "sqlite://file::root:?cache=shared&mode=memory",
	}
	router := api.CreateRouter(config)

	get(t, router, schemas.UserListSchema{Users: []schemas.UserSchema{}})
	post(t, router, "test")
	get(t, router, schemas.UserListSchema{Users: []schemas.UserSchema{{ID: 1, Username: "test"}}})
}

func get(t *testing.T, router *fiber.App, expected schemas.UserListSchema) {
	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)

	AssertJsonResponse(t, resp, expected)
}

func post(t *testing.T, router *fiber.App, username string) {
	req := httptest.NewRequest("POST", "/users", strings.NewReader(`{"username": "`+username+`"}`))
	resp, _ := router.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
	AssertJsonResponse(t, resp, schemas.UserSchema{ID: 1, Username: username})
}
