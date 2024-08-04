package middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/pkg/ctxext"
	"github.com/rs/zerolog"
)

func SetLogger(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Add fields to the logger
			rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
			logger = logger.With().Str("request_id", rID).Logger()

			// Attach the logger to the context
			ctx = ctxext.SetLogger(ctx, logger)
			return next(ctx)
		}
	}
}

func LogRequest(logger zerolog.Logger) echo.MiddlewareFunc {
	return echomw.RequestLoggerWithConfig(echomw.RequestLoggerConfig{
		LogRequestID:     true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogReferer:       true,
		LogProtocol:      true,
		LogStatus:        true,
		LogLatency:       true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogUserAgent:     true,
		LogError:         true,
		LogValuesFunc: func(ctx echo.Context, v echomw.RequestLoggerValues) error {
			// Log level
			var level zerolog.Level
			switch {
			case v.Status >= 500:
				level = zerolog.ErrorLevel
			case v.Status >= 400:
				level = zerolog.WarnLevel
			default:
				level = zerolog.InfoLevel
			}

			// User
			userID := ""
			oryID := ""
			if user, ok := ctxext.GetAuthUser(ctx); ok {
				userID = strconv.Itoa(user.ID)
			}
			if ory, ok := ctxext.GetOryIdentity(ctx); ok {
				oryID = ory.GetOryID().String()
			}

			logger.WithLevel(level).
				Str("request_id", v.RequestID).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("method", v.Method).
				Str("uri", v.URI).
				Str("referer", v.Referer).
				Str("protocol", v.Protocol).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Str("latency_human", v.Latency.String()).
				Str("request_size", v.ContentLength).
				Int64("response_size", v.ResponseSize).
				Str("user_agent", v.UserAgent).
				Str("user", userID).
				Str("ory_id", oryID).
				Err(v.Error).
				Msg("request")

			return nil
		},
	})
}
