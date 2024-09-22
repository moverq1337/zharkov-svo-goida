package main

import (
	"github.com/go-playground/validator/v10"
	authLib "github.com/neiasit/auth-library"
	grpcLib "github.com/neiasit/grpc-library"
	httpSupport "github.com/neiasit/http-support-library"
	loggingLib "github.com/neiasit/logging-library"
	redisLib "github.com/neiasit/redis-library"
	"github.com/neiasit/service-boilerplate/internal/example_domain_first"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	app := fx.New(

		// setting validator
		fx.Provide(func() *validator.Validate {
			return validator.New(
				validator.WithRequiredStructEnabled(),
			)
		}),

		// including platform libs here
		loggingLib.Module,
		grpcLib.ModuleWithAuth,
		httpSupport.Module,
		redisLib.Module,
		authLib.AuthKeycloakModule,

		// setting logger
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),

		// including app modules here
		example_domain_first.Module,
	)

	app.Run()
}
