package http

import (
	"github.com/gofiber/fiber"
	"github.com/kumarabd/go-http/pkg/app"
	"github.com/kumarabd/gokit/models"
)

func Health(c *fiber.Ctx) {
	c.Status(200)
	c.JSON(&models.Response{
		Code: "200",
		Body: app.GetHealth(),
	})
}

func (s *Service) Stats(c *fiber.Ctx) {
	c.Status(200)
	c.JSON(&models.Response{
		Code: "200",
		Body: app.GetStats(&app.Stats{
			Name:      s.Name,
			Port:      s.Port,
			Proxy:     s.Proxy,
			Version:   s.Version,
			StartedAt: s.StartedAt,
		}),
	})
}
