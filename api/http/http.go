package http

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/kumarabd/gokit/logger"
)

type Service struct {
	Name      string    `json:"name"`
	Config    string    `json:"config"`
	Port      string    `json:"port"`
	Proxy     string    `json:"proxy"`
	Version   string    `json:"version"`
	StartedAt time.Time `json:"startedat,string"`
}

func RecoveryHandler() func(*fiber.Ctx) {
	return recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			logger.Log(fmt.Sprintf("%v for request, Request: %v,", err, c.Fasthttp.Request))
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	})
}

func Start(s *Service) error {

	// Init application handler
	app := fiber.New()
	app.Use(RecoveryHandler())

	// Init routes
	app.Get("/health", Health)

	// Init server
	err := app.Listen(s.Port)
	if err != nil {
		return err
	}
	return nil
}
