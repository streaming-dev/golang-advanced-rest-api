package middleware

import (
	"github.com/AleksK1NG/api-mc/pkg/logger"
	"github.com/labstack/echo/v4"
	"time"
)

// Request logger middleware
func RequestLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			start := time.Now()
			err := next(ctx)

			req := ctx.Request()
			res := ctx.Response()
			status := res.Status
			size := res.Size
			s := time.Since(start).String()

			logger.Infof("TimeSince: %s, Method: %s, URI: %s, Status: %v, Size: %v", s, req.Method, req.URL, status, size)
			return err
		}
	}
}
