package httpserver

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/professionsforall/hexagonal-template/pkg/log/logger"
)

var HttpServer *httpServer

type httpServer struct {
	app    *fiber.App
	logger logger.AppLogger
	port   string
}

func Apply(app *fiber.App, port string, logger logger.AppLogger) {
	HttpServer = &httpServer{
		app:    app,
		logger: logger,
		port:   port,
	}
}

func (h *httpServer) Add(prefix string, fiberApp *fiber.App) {
	h.app.Mount(prefix, fiberApp)
}

func (h *httpServer) Start() error {
	return h.app.Listen(":" + h.port)
}

func (h *httpServer) ShutDown(ctx context.Context) error {
	return h.app.ShutdownWithContext(ctx)
}
