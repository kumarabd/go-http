package http

import (
	"github.com/gofiber/fiber"
	"github.com/kumarabd/go-backend/pkg/health"
	"github.com/kumarabd/gokit/models"
)

func Health(c *fiber.Ctx) {
	c.Status(200)
	c.Send(&models.Response{
		Code: "200",
		Body: health.GetHealth(),
	})
}
