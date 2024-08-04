package middleware

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/pkg/ctxext"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestLogRequestID(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")

	out := &bytes.Buffer{}
	logger := zerolog.New(out)
	ctxext.SetLogger(ctx, logger)

	require.NoError(t, tests.ExecuteMiddleware(ctx, echomw.RequestID()))
	require.NoError(t, tests.ExecuteMiddleware(ctx, SetLogger(logger)))
	// require.NoError(t, tests.ExecuteMiddleware(ctx, LogRequest(log.Logger)))

	ctxext.Logger(ctx).Info().Msg("test")
	rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	assert.Contains(t, out.String(), fmt.Sprintf(`id":"%s"`, rID))
}

func TestLogRequest(t *testing.T) {
	statusCode := 200
	out := &bytes.Buffer{}

	exec := func() {
		ctx, _ := tests.NewContext(c.Web, "http://test.localhost/abc?d=1&e=2")
		logger := zerolog.New(out).With().Str("previous", "param").Logger()
		ctxext.SetLogger(ctx, logger)
		ctx.Request().Header.Set("Referer", "ref.com")
		ctx.Request().Header.Set(echo.HeaderXRealIP, "21.12.12.21")

		require.NoError(t, tests.ExecuteHandler(ctx, func(ctx echo.Context) error {
			return ctx.String(statusCode, "hello")
		},
			SetLogger(logger),
			LogRequest(logger),
		))
	}

	exec()
	assert.Contains(t, out.String(), `"previous":"param"`)
	assert.Contains(t, out.String(), `"remote_ip":"21.12.12.21"`)
	assert.Contains(t, out.String(), `"host":"test.localhost"`)
	assert.Contains(t, out.String(), `"referer":"ref.com"`)
	assert.Contains(t, out.String(), `"status":"200"`)
	assert.Contains(t, out.String(), `"request_size":"0"`)
	assert.Contains(t, out.String(), `"response_size":"5"`)
	assert.Contains(t, out.String(), `"latency":"1"`)
	assert.Contains(t, out.String(), `"level":"INFO"`)
	assert.Contains(t, out.String(), `"method":"GET"`)
	assert.Contains(t, out.String(), `"uri":"/abc?d=1&e=2"`)

	statusCode = 500
	exec()
	assert.Contains(t, out.String(), `"level":"ERROR"`)
}
