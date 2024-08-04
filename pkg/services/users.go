package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
)

type (
	UsersClient struct {
		orm *ent.Client
	}

	UserCreateData struct {
		OryID      uuid.UUID
		UILanguage string
	}

	UserUpdateData struct {
		ID         int
		UILanguage string
	}
)

func NewUsersClient(orm *ent.Client) *UsersClient {
	return &UsersClient{orm: orm}
}

func (c *UsersClient) Get(ctx echo.Context, oryID uuid.UUID) (*ent.User, error) {
	return c.orm.User.Query().Where(user.OryID(oryID)).Only(ctx.Request().Context())
}

func (c *UsersClient) Create(ctx echo.Context, data UserCreateData) (*ent.User, error) {
	tx, err := c.orm.Tx(ctx.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("Error starting a transaction: %w", err)
	}

	user, err := tx.User.
		Create().
		SetOryID(data.OryID).
		SetUILanguage(user.UILanguage(data.UILanguage)).
		Save(ctx.Request().Context())
	if err != nil {
		return nil, rollback(tx, err, "Error creating user")
	}

	return user, tx.Commit()
}

func (c *UsersClient) Update(ctx echo.Context, data UserUpdateData) (*ent.User, error) {
	return c.orm.User.
		UpdateOneID(data.ID).
		SetNillableUILanguage((*user.UILanguage)(&data.UILanguage)).
		Save(ctx.Request().Context())
}
