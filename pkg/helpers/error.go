package helpers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator"
)

type Error struct {
	ID      string      `json:"id" xml:"id"`
	Details interface{} `json:"details" xml:"details"`
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Details == nil {
		return fmt.Sprintf("Error: [%s]", e.ID)
	}
	return fmt.Sprintf("Error: [%s] %+v", e.ID, e.Details)
}

func NewError(id string) *Error {
	return &Error{ID: id}
}

func NewErrorDetails(id string, details interface{}) *Error {
	return &Error{ID: id, Details: details}
}

func NewErrorFromError(err error) *Error {
	if err == nil {
		return nil
	}
	if pbErr, ok := err.(*Error); ok {
		return pbErr
	}
	jsonUnmarshalTypeError, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return &Error{ID: "json_type_error", Details: formatJSONUnmarshalTypeError(jsonUnmarshalTypeError)}
	}
	jsonSyntaxError, ok := err.(*json.SyntaxError)
	if ok {
		return &Error{ID: "json_syntax_error", Details: map[string]interface{}{"position": jsonSyntaxError.Offset, "error": jsonSyntaxError.Error()}}
	}
	xmlSyntaxError, ok := err.(*xml.SyntaxError)
	if ok {
		return &Error{ID: "xml_syntax_error", Details: map[string]interface{}{"line": xmlSyntaxError.Line, "error": xmlSyntaxError.Msg}}
	}
	if err == io.ErrUnexpectedEOF {
		return &Error{ID: "invalid_indata", Details: "Input too short"}
	}
	if validatorErr, ok := err.(validator.ValidationErrors); ok {
		return &Error{ID: "validation_error", Details: formatValidationErrors(validatorErr)}
	}
	return NewErrorDetails("internal_server_error", err.Error())
}

func formatValidationErrors(err validator.ValidationErrors) []map[string]interface{} {
	v := make([]map[string]interface{}, 0)
	for _, e := range err {
		splits := strings.SplitN(e.Namespace(), ".", 2)
		v = append(v, map[string]interface{}{
			"field":           splits[1],
			"struct":          splits[0],
			"type":            e.Kind().String(),
			"validation":      e.Tag(),
			"validationParam": e.Param(),
			"value":           e.Value(),
		})
	}
	return v
}

func formatJSONUnmarshalTypeError(err *json.UnmarshalTypeError) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"field":    err.Field,
			"expected": err.Type.Kind().String(),
			"actual":   err.Value,
		},
	}
}
