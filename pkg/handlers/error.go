package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/ctxext"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

type Error struct {
	*services.TemplateRenderer
}

func (e *Error) Page(err error, ctx echo.Context) {
	if ctx.Response().Committed || ctxext.IsCanceledError(err) {
		return
	}

	// Determine the error status code
	code := http.StatusInternalServerError
	var httpError *echo.HTTPError
	var notFoundError *ent.NotFoundError
	switch {
	// case errors.Is(err, privacyrules.ErrDenyUnauthorized):
	// 	code = http.StatusUnauthorized
	// case errors.Is(err, privacyrules.ErrDenyForbidden):
	// 	code = http.StatusForbidden
	case errors.As(err, &httpError):
		code = httpError.Code
	case errors.As(err, &notFoundError):
		code = http.StatusNotFound
	}

	// Send the response
	page, ok := ctxext.IsPageRoute(ctx)
	if (ok && !page) || (!ok &&
		(strings.Contains(ctx.Request().Header.Get("Accept"), "json")) ||
		(strings.Contains(ctx.Request().Header.Get("Content-Type"), "json"))) {
		e.apiErrorHandler(ctx, err, code)
	} else {
		e.pageErrorHandler(ctx, code)
	}
}

func (e *Error) pageErrorHandler(ctx echo.Context, code int) {
	page := page.New(ctx)
	page.Title = http.StatusText(code)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageError
	page.StatusCode = code
	page.HTMX.Request.Enabled = false

	if err := e.RenderPage(ctx, page); err != nil {
		ctxext.Logger(ctx).Error().Err(err).Msg("Error sending page error response")
	}
}

func (e *Error) apiErrorHandler(ctx echo.Context, err error, code int) {
	if err := ctx.JSON(code, err); err != nil {
		ctxext.Logger(ctx).Error().Err(err).Msg("Error sending API error response")
	}
}
