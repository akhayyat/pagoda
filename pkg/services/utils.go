package services

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
)

func rollback(tx *ent.Tx, err error, msg string) error {
	err = fmt.Errorf("%v: %w", msg, err)
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func internalRedirect(ctx echo.Context, path string) {
	ctx.Response().Header().Set("X-Accel-Redirect", path)
	ctx.Response().Header().Set("X-Accel-Buffering", "no")
}
