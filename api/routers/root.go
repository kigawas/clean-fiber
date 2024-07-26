package routers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/clean-fiber/app"
)

func CreateRootRouter(router fiber.Router) {
	router.Get("/", rootGet)
}

// rootGet godoc
//
//	@Summary		Health check
//	@Description	Execute hello world SQL in DB
//	@Tags			root
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Router			/ [get]
func rootGet(c fiber.Ctx) error {
	var result string
	app.GetDB().Raw("SELECT 'Hello, World from DB!'").Scan(&result)
	return c.SendString(result)
}
