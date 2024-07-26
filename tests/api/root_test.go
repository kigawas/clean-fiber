package tests

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/kigawas/clean-fiber/api"
	"github.com/kigawas/clean-fiber/app"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	config := app.Config{
		DatabaseURL: "sqlite://file::root:?cache=shared&mode=memory",
	}
	router := api.CreateRouter(config)

	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := router.Test(req)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, World from DB!", string(body))
}
