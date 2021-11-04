package goladok3

import (
	"context"
	"fmt"
	"strings"
)

// IsLadokPermissionsSufficient compare ladok permissions with ps
func (c *Client) IsLadokPermissionsSufficient(ctx context.Context, myPermissions Permissions) (bool, error) {
	var (
		e             = &Errors{}
		internalError = []InternalError{}
	)

	egna, _, err := c.Kataloginformation.GetAnvandarbehorighetEgna(ctx)
	if err != nil {
		return false, err
	}

	if len(egna.Anvandarbehorighet) < 1 {
		return false, ErrNotSufficientPermissions
	}

	ladokProfile, _, err := c.Kataloginformation.GetBehorighetsprofil(ctx, &GetBehorighetsprofilerCfg{UID: egna.UID})
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
			internalError = append(internalError, InternalError{
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
	case envIntTestAPI:
		return envIntTestAPI, nil
	case envProdAPI:
		return envProdAPI, nil
	case envTestAPI:
		return envTestAPI, nil
	default:
		return "", ErrNoEnvFound
	}
}

// trim remove "urn"
func (id FeedID) trim() string {
	return strings.Split(string(id), ":")[2]
}
