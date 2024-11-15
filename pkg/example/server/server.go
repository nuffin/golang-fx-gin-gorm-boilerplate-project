package server

import (
	"boilerplate/config"
	"boilerplate/internal/utils"
	"boilerplate/internal/web/server"
	"boilerplate/pkg/example/server/handlers"
	"boilerplate/pkg/example/server/services"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ExampleAppServer struct {
	server.Server
}

func NewExampleAppServer(c *config.Config, d *gorm.DB, l *zap.Logger) *ExampleAppServer {
	s := &ExampleAppServer{
		Server: *server.NewServer(&c.Config, d, l),
	}
	s.ConfigureRouteGroups()
	return s
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewExampleAppServer),
	),

	services.Module,
	handlers.Module,

	fx.Invoke(
		func(s *ExampleAppServer, logger *zap.Logger) {
			logger.Debug("ExampleAppServer module invoked")
			go func() {
				_ = s.Run(utils.GetWebserverAddr())
			}()
		},
		func(s *ExampleAppServer, handler *handlers.PingHandler) {
			s.Handlers = append(s.Handlers, handler)
			s.NoAuth.GET("/ping", handler.Ping())
		},
	),
)
