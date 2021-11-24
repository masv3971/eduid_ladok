package goladok3

import (
	"errors"
	"fmt"

	"github.com/masv3971/goladok3/ladoktypes"
)

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
	// ErrNoEnvFound if no valid environment is found in certificate (ou)
	ErrNoEnvFound = &Errors{Internal: []ladoktypes.InternalError{{Msg: "No valid ladok-environment (OU) found in certificate"}}}
	// ErrNotSufficientPermissions if not all provided permissions are met
	ErrNotSufficientPermissions = &Errors{Internal: []ladoktypes.InternalError{{Msg: "No permissions found in ladok", Type: "Permission"}}}
	// ErrNoPermissionProvided when input Permission is empty
	ErrNoPermissionProvided = &Errors{Internal: []ladoktypes.InternalError{{Msg: "No permissions provided", Type: "Permission"}}}
)

// Errors is the bespoke error struct
type Errors struct {
	Internal []ladoktypes.InternalError `json:"details,omitempty"`
	Ladok    *ladoktypes.LadokError     `json:"ladok,omitempty"`
}

func (e *Errors) Error() string {
	if e.Ladok != nil && len(e.Internal) > 0 {
		return fmt.Sprintf("internal error: %v, ladok error: %v", e.Internal, e.Ladok)
	} else if len(e.Internal) > 0 {
		return fmt.Sprintf("internal error: %v", e.Internal)
	} else if e.Ladok != nil {
		return fmt.Sprintf("ladok error: %v", e.Ladok)
	}
	return ""
}

// Error interface
type Error interface {
	Error() string
}

func oneError(m, t, f, e string) *Errors {
	return &Errors{Internal: []ladoktypes.InternalError{{Msg: m, Type: t, Func: f, PreviousError: e}}}
}
