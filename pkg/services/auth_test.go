package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestAuthClient_Auth(t *testing.T) {
	assertNoAuth := func() {
		_, _, err := c.Auth.GetAuthenticatedUser(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
	}

	assertNoAuth()

	// err := c.Auth.Login(ctx, usr.ID)
	// require.NoError(t, err)

	u, o, err := c.Auth.GetAuthenticatedUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, u.ID, usr[0].ID)
	assert.Equal(t, o.OryID, oid[0].OryID)

	// err = c.Auth.Logout(ctx)
	// require.NoError(t, err)

	assertNoAuth()
}
