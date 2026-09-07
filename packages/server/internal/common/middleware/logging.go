package middleware

import (
	"my-codelens-app/internal/common/logger"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
)

func ZapLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,

		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			fields := []zap.Field{
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
			}

			if v.Error != nil {
				fields = append(fields,
					zap.String("error", v.Error.Error()),
				)
			}

			logger.Info("Incoming request", fields...)
			return nil
		},
	})
}

func Recovery() echo.MiddlewareFunc {
	return middleware.Recover()
}
