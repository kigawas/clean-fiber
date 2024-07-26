package tests

import (
	"testing"

	"github.com/kigawas/clean-fiber/app"
	"github.com/kigawas/clean-fiber/app/persistence"
	"github.com/kigawas/clean-fiber/models/params"
	"github.com/kigawas/clean-fiber/models/queries"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	db := app.SetupDB("sqlite://file::user:?cache=shared&mode=memory", &gorm.Config{})
	app.MigrateDB(db)

	user, _ := persistence.CreateUser(db, &params.CreateUserParams{Username: "test"})
	assert.Equal(t, "test", user.Username)

	user, _ = persistence.GetUser(db, 1)
	assert.Equal(t, "test", user.Username)

	users, _ := persistence.GetUsers(db, &queries.UserQuery{Username: "test"})
	assert.Equal(t, "test", users.Users[0].Username)
}
