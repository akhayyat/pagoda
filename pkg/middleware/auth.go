package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/ctxext"
	"github.com/mikestefanello/pagoda/pkg/services"
)

func LoadAuthenticatedUser(auth *services.OryAuthClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, oryIdentity, err := auth.GetAuthenticatedUser(ctx)
			// TODO: Replace err.(type) with erros.Is and errors.As
			switch err.(type) {
			case nil: // Authenticated and recognized
				ctx.Set(ctxext.OryIdentityKey, oryIdentity)
				ctx.Set(ctxext.AuthUserKey, user)
				// Make AuthUserKey available in ctx.Request().Context()
				ctx.SetRequest(ctx.Request().WithContext(
					context.WithValue(
						ctx.Request().Context(), ctxext.AuthUserKey, user)))
			case *ent.NotFoundError: // Authenticated but not created
				ctx.Set(ctxext.OryIdentityKey, oryIdentity)
			case services.NotAuthenticatedError:
				ctxext.Logger(ctx).Debug().Err(err).Msg("Unauthenticated")
			default:
				ctxext.Logger(ctx).Error().Err(err).Msg("Authentication failed: unexpected error")
			}
			return next(ctx)
		}
	}
}

func isAuthButNoUser(ctx echo.Context) bool {
	_, oryOK := ctxext.GetOryIdentity(ctx)
	_, userOK := ctxext.GetAuthUser(ctx)
	return oryOK && !userOK
}

func forbidden(ctx echo.Context) error {
	if page, _ := ctxext.IsPageRoute(ctx); page {
		return ctx.Redirect(http.StatusSeeOther, ctx.Echo().Reverse("page-user-new"))
	} else { // API
		return echo.NewHTTPError(
			http.StatusForbidden,
			fmt.Sprintf("No user found for this identity. Please create a user at %v",
				ctx.Echo().Reverse("api-user-create")),
		)
	}
}

// Auth optional: If the user is authenticated, require they have a user account
func AuthOptional() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if isAuthButNoUser(ctx) {
				return forbidden(ctx)
			}
			return next(ctx)
		}
	}
}

func hasAuthUser(ctx echo.Context) bool {
	_, oryOK := ctxext.GetOryIdentity(ctx)
	_, userOK := ctxext.GetAuthUser(ctx)
	return oryOK && userOK
}

func unauthorized(ctx echo.Context) error {
	if page, _ := ctxext.IsPageRoute(ctx); page {
		return ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
}

// Require both auth and user (AND)
func AuthRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if !hasAuthUser(ctx) {
				return unauthorized(ctx)
			}
			return next(ctx)
		}
	}
}

// Ensure user is Ory-authenticated but does not have a user account
func AuthIncomplete() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if _, ok := ctxext.GetOryIdentity(ctx); !ok {
				return unauthorized(ctx)
			} else if _, ok := ctxext.GetAuthUser(ctx); ok {
				if page, _ := ctxext.IsPageRoute(ctx); page {
					return redirectToNext(ctx)
				} else {
					return echo.NewHTTPError(http.StatusBadRequest,
						"This identity already has a user")
				}
			}
			return next(ctx)
		}
	}
}

func redirectToNext(ctx echo.Context) error {
	var nextURL string
	err := echo.QueryParamsBinder(ctx).String("next", &nextURL).BindError()
	if err != nil || nextURL == "" {
		nextURL = ctx.Echo().Reverse("home")
	}
	return ctx.Redirect(http.StatusSeeOther, nextURL)
}
