package ctxext

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/rs/zerolog"
)

// OryIdentity interface makes the services.OryIdentity structure accessible
// without import cycles
type OryIdentity interface {
	GetOryID() uuid.UUID
	GetEmail() string
	GetVerified() bool
}

const (
	// IsPageRouteKey is the key used tp specify whether the current route is a
	// page route (true) or an API route (false)
	IsPageRouteKey = "is_page_route"

	// AuthUserKey is the key value used to store the authenticated user in context
	AuthUserKey = "auth_user"

	// OryIdentityKey is the key value used to store Ory-authenticated identity,
	// regardless of its existence in Tfarraj
	OryIdentityKey = "ory_identity"

	// // UserKey is the key value used to store a user in context
	// UserKey = "user"

	// FormKey is the key value used to store a form in context
	FormKey = "form"

	// SessionKey is the key value used to store the session data in context
	SessionKey = "session"
)

// IsCanceledError determines if an error is due to a context cancelation
func IsCanceledError(err error) bool {
	return errors.Is(err, context.Canceled)
}

func SetLogger(ctx echo.Context, logger zerolog.Logger) echo.Context {
	stdCtx := logger.WithContext(ctx.Request().Context())
	ctx.SetRequest(ctx.Request().WithContext(stdCtx))
	return ctx
}

func Logger(ctx echo.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx.Request().Context())
}

//////////////////////////////////////////////////
//  Typed Accessors
//////////////////////////////////////////////////

func GetOryIdentity(ctx echo.Context) (OryIdentity, bool) {
	i, ok := ctx.Get(OryIdentityKey).(OryIdentity)
	return i, ok
}

func GetAuthUser(ctx echo.Context) (*ent.User, bool) {
	u, ok := ctx.Get(AuthUserKey).(*ent.User)
	return u, ok
}

func IsPageRoute(ctx echo.Context) (bool, bool) {
	p, ok := ctx.Get(IsPageRouteKey).(bool)
	return p, ok
}
