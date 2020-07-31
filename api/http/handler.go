package http

import (
	"github.com/gofiber/fiber"
	"github.com/kumarabd/go-http/pkg/platform"
	"github.com/kumarabd/gokit/models"
)

func Health(c *fiber.Ctx) {
	c.Status(200)
	c.JSON(&models.Response{
		Code: "200",
		Body: platform.GetHealth(),
	})
}

func (s *Service) Stats(c *fiber.Ctx) {
	c.Status(200)
	c.JSON(&models.Response{
		Code: "200",
		Body: platform.GetStats(&models.Stats{
			Name:      s.Name,
			Port:      s.Port,
			Proxy:     s.Proxy,
			Version:   s.Version,
			StartedAt: s.StartedAt,
		}),
	})
}
