package routers

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	"github.com/kigawas/clean-fiber/api/models"
	"github.com/kigawas/clean-fiber/app"
	"github.com/kigawas/clean-fiber/app/persistence"
	"github.com/kigawas/clean-fiber/models/params"
	"github.com/kigawas/clean-fiber/models/queries"
	"gorm.io/gorm"
)

func CreateUserRouter(router fiber.Router) {
	router.Get("/", usersGet).Post("/", userPost).Get("/:id", userIdGet)
}

// usersGet godoc
//
//	@Summary		Show users
//	@Description	get users by username
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	false	"Username"
//	@Success		200	{object}	schemas.UserListSchema
//	@Failure		400	{object}	models.ApiErrorResponse
//	@Failure		404	{object}	models.ApiErrorResponse
//	@Failure		500	{object}	models.ApiErrorResponse
//	@Router			/users [get]
func usersGet(c fiber.Ctx) error {
	query := &queries.UserQuery{}
	if err := c.Bind().Query(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.FromString("invalid query"))
	}

	users, err := persistence.GetUsers(app.GetDB(), query)
	if err != nil {
		return errors.New("db error") // hide details
	}
	return c.JSON(users)
}

// userIdGet godoc
//
//	@Summary		Show a user
//	@Description	get user by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	schemas.UserSchema
//	@Failure		400	{object}	models.ApiErrorResponse
//	@Failure		404	{object}	models.ApiErrorResponse
//	@Failure		500	{object}	models.ApiErrorResponse
//	@Router			/users/{id} [get]
func userIdGet(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	user, err := persistence.GetUser(app.GetDB(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.FromString("user not found"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.FromString("db error"))
	}
	return c.JSON(user)
}

// userPost godoc
//
//	@Summary		Create a user
//	@Description	Create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body	params.CreateUserParams	true	"User params"
//	@Success		201	{object}	schemas.UserSchema
//	@Failure		400	{object}	models.ApiErrorResponse
//	@Failure		500	{object}	models.ApiErrorResponse
//	@Router			/users [post]
func userPost(c fiber.Ctx) error {
	payload := &params.CreateUserParams{}
	if err := c.Bind().JSON(payload); err != nil {
		return err
	}

	user, err := persistence.CreateUser(app.GetDB(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.FromString("db error"))
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
