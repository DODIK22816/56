package servers

import (
	"context"

	"github.com/DODIK22816/final5/internal/orchestrator/delivery/rest/routes"
	"github.com/DODIK22816/final5/internal/orchestrator/service"

	"github.com/labstack/echo/v4"
)

type CalculatorServerConfig struct {
	Port string `env:"PORT" env-default:"8080"`
}

type CalculatorServer struct {
	cfg    CalculatorServerConfig
	engine *echo.Echo
}

func NewCalculatorServer(cfg CalculatorServerConfig, service service.CalculatorService) *CalculatorServer {
	e := echo.New()
	routes.CalculatorRoutes(e, service)
	return &CalculatorServer{cfg: cfg, engine: e}
}

func (s *CalculatorServer) Start() error {
	return s.engine.Start(":" + s.cfg.Port)
}

func (s *CalculatorServer) Stop(ctx context.Context) error {
	return s.engine.Shutdown(ctx)
}
