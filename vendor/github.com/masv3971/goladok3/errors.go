package goladok3

import (
	"errors"
	"fmt"
)

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
	// ErrNoEnvFound if no valid environment is found in certificate (ou)
	ErrNoEnvFound = &Errors{Internal: []InternalError{{Msg: "No valid ladok-environment (OU) found in certificate"}}}
	// ErrNotSufficientPermissions if not all provided permissions are met
	ErrNotSufficientPermissions = &Errors{Internal: []InternalError{{Msg: "No permissions found in ladok", Type: "Permission"}}}
	// ErrNoPermissionProvided when input Permission is empty
	ErrNoPermissionProvided = &Errors{Internal: []InternalError{{Msg: "No permissions provided", Type: "Permission"}}}
)

// InternalError type
type InternalError struct {
	Msg           string `json:"msg"`
	Type          string `json:"type"`
	Func          string `json:"func"`
	PreviousError string `json:"previous_error"`
}

// LadokError returns by Ladok
type LadokError struct {
	Detaljkod       string        `json:"Detaljkod"`
	DetaljkodText   string        `json:"DetaljkodText"`
	FelUID          string        `json:"FelUID"`
	Felgrupp        string        `json:"Felgrupp"`
	FelgruppText    string        `json:"FelgruppText"`
	Felkategori     string        `json:"Felkategori"`
	FelkategoriText string        `json:"FelkategoriText"`
	Meddelande      string        `json:"Meddelande"`
	Link            []interface{} `json:"link"`
}

// Errors is the bespoke error struct
type Errors struct {
	Internal []InternalError `json:"details,omitempty"`
	Ladok    *LadokError     `json:"ladok,omitempty"`
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
	return &Errors{Internal: []InternalError{{Msg: m, Type: t, Func: f, PreviousError: e}}}
}
