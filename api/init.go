package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/kigawas/clean-fiber/api/routers"
	"github.com/kigawas/clean-fiber/api/ws"
	"github.com/kigawas/clean-fiber/app"
)

func createConfig() fiber.Config {
	return fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	}
}

// TODO: revamp error handling
func setupRouter(fApp *fiber.App) {
	var m = map[string]func(fiber.Router){
		"/":      routers.CreateRootRouter,
		"/users": routers.CreateUserRouter,
	}
	for path, createRouter := range m {
		createRouter(fApp.Group(path))
	}
}

func CreateRouter(config app.Config) *fiber.App {
	db := app.SetupDB(config.DatabaseURL, &gorm.Config{})
	app.MigrateDB(db)

	router := fiber.New(createConfig())
	setupRouter(router)
	ws.SetupWS(router)
	return router
}
