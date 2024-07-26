package ws

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kigawas/clean-fiber/websocket"
)

func SetupWS(router fiber.Router) {
	router.Use("/ws", func(c fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		for {
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("read error, exiting:", err)
				break
			}

			resp := Response{
				ID:      id,
				Content: msg.Content,
			}
			if err = c.WriteJSON(resp); err != nil {
				log.Println("write error, exiting:", err)
				break
			}
		}
	}))
}
