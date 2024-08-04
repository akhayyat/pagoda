package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/ctxext"
	ory "github.com/ory/client-go"
)

type OryIdentity struct {
	OryID    uuid.UUID
	Email    string
	Verified bool
}

func (i *OryIdentity) GetOryID() uuid.UUID {
	return i.OryID
}

func (i *OryIdentity) GetEmail() string {
	return i.Email
}

func (i *OryIdentity) GetVerified() bool {
	return i.Verified
}

//////////////////////////////////////////////////
//  OryAuthClient
//////////////////////////////////////////////////

type OryAuthClient struct {
	ory   *ory.APIClient
	users *UsersClient
}

type FlowAttrs struct {
	Action        string
	CSRF          string
	FieldMessages map[string][]ory.UiText
	Messages      []ory.UiText
	State         interface{}
}

// NotAuthenticatedError is an error returned when a user is not authenticated
type NotAuthenticatedError struct{}

// Error implements the error interface.
func (e NotAuthenticatedError) Error() string {
	return "user not authenticated by Ory"
}

func newOryAuthClient(cfg *config.Config, users *UsersClient) *OryAuthClient {
	oryConf := ory.NewConfiguration()
	oryConf.Servers = ory.ServerConfigurations{{URL: cfg.Ory.URL}}
	return &OryAuthClient{
		ory:   ory.NewAPIClient(oryConf),
		users: users,
	}
}

func (c *OryAuthClient) getOryIdentity(ctx echo.Context) (*OryIdentity, error) {
	cookie, _ := ctx.Request().Cookie("ory_kratos_session")
	authFields := strings.Fields(ctx.Request().Header.Get("Authorization"))
	req := c.ory.FrontendAPI.ToSession(ctx.Request().Context())
	if cookie != nil {
		req = req.Cookie(cookie.String())
	}
	if len(authFields) == 2 && strings.ToLower(authFields[0]) == "bearer" {
		req = req.XSessionToken(authFields[1])
	}
	session, _, err := req.Execute()

	if (err != nil && session == nil) || (err == nil && !*session.Active) {
		return nil, NotAuthenticatedError{}
	} else {
		// Extract OryIdentity struct from `session`
		oryID, err := uuid.Parse(session.Identity.Id)
		if err != nil {
			ctxext.Logger(ctx).Error().Str("ory_id", session.Identity.Id).Err(err).
				Msg("Received OryID that is not a valid UUID")
			return nil, NotAuthenticatedError{}
		}
		email := session.Identity.Traits.(map[string]interface{})["email"].(string)
		verified := false
		for _, address := range session.Identity.VerifiableAddresses {
			if address.Value == email {
				verified = address.Verified
				break
			}
		}

		return &OryIdentity{
			OryID:    oryID,
			Email:    email,
			Verified: verified,
		}, nil
	}
}

func (c *OryAuthClient) GetAuthenticatedUser(ctx echo.Context) (*ent.User, *OryIdentity, error) {
	if identity, err := c.getOryIdentity(ctx); err == nil {
		if user, err := c.users.Get(ctx, identity.OryID); err == nil {
			return user, identity, nil
		} else {
			return nil, identity, err
		}
	} else {
		return nil, identity, err
	}
}

func (c *OryAuthClient) GetEmails(ctx context.Context, oryIDs []string) (map[string]string, error) {
	identities, _, err := c.ory.IdentityAPI.ListIdentities(ctx).Ids(oryIDs).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving Ory identities: %w", err)
	}
	emails := make(map[string]string, len(oryIDs))
	for _, identity := range identities {
		emails[identity.Id] = identity.Traits.(map[string]interface{})["email"].(string)
	}
	return emails, nil
}
