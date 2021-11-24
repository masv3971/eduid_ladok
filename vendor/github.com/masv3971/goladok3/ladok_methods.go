package goladok3

import (
	"context"
	"fmt"

	"github.com/masv3971/goladok3/ladoktypes"
)

// IsLadokPermissionsSufficient compare ladok permissions with ps
func (c *Client) IsLadokPermissionsSufficient(ctx context.Context, myPermissions Permissions) (bool, error) {
	var (
		e             = &Errors{}
		internalError = []ladoktypes.InternalError{}
	)

	egna, _, err := c.Kataloginformation.GetAnvandarbehorighetEgna(ctx)
	if err != nil {
		return false, err
	}

	if len(egna.Anvandarbehorighet) < 1 {
		return false, ErrNotSufficientPermissions
	}

	ladokProfile, _, err := c.Kataloginformation.GetBehorighetsprofil(ctx, &GetBehorighetsprofilerReq{UID: egna.UID})
	if err != nil {
		return false, err
	}

	if len(ladokProfile.Systemaktiviteter) == 0 {
		return false, ErrNotSufficientPermissions
	}
	if len(myPermissions) == 0 {
		return false, ErrNoPermissionProvided
	}

	for myPermissionsID, myPermissionsValue := range myPermissions {
		notFound := true
		for _, systemaktivitet := range ladokProfile.Systemaktiviteter {
			if systemaktivitet.ID == myPermissionsID {
				if myPermissionsValue == systemaktivitet.Rattighetsniva {
					notFound = false
					continue
				}
			}
		}
		if notFound {
			//missingPermission[myPermissionsID] = myPermissionsValue
			internalError = append(internalError, ladoktypes.InternalError{
				Msg:  fmt.Sprintf("Missing id: %d, value: %q", myPermissionsID, myPermissionsValue),
				Type: "Permission",
			})
		}
	}
	if len(internalError) > 0 {
		e.Internal = internalError
		return false, e

	}
	return true, nil
}

func (c *Client) environment() (string, error) {
	switch c.certificate.Subject.OrganizationalUnit[1] {
	case ladoktypes.EnvIntTestAPI:
		return ladoktypes.EnvIntTestAPI, nil
	case ladoktypes.EnvProdAPI:
		return ladoktypes.EnvProdAPI, nil
	case ladoktypes.EnvTestAPI:
		return ladoktypes.EnvTestAPI, nil
	default:
		return "", ErrNoEnvFound
	}
}
