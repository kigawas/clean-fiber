package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/clean-fiber/api"
	"github.com/kigawas/clean-fiber/app"
)

// @title Clean Fiber
// @version 1.0
// @description Fiber with clean architecture
// @license.name MIT
// @license.url https://opensource.org/license/mit
// @host localhost:3001
// @BasePath /
func main() {
	app.LoadEnv()
	config := app.FromEnv()
	router := api.CreateRouter(config)
	defer router.Shutdown()

	if err := router.Listen(config.URL(), fiber.ListenConfig{
		EnablePrefork: config.Prefork,
	}); err != nil {
		log.Fatal(err)
	}
}
