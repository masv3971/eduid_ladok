package helpers

import (
	"eduid_ladok/pkg/logger"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

// Check checks for validation error
func Check(s interface{}, log *logger.Logger) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("ERR: Field %q of type %q violates rule: %q\n", err.Namespace(), err.Kind(), err.Tag())
			log.Error(msg)
		}
		return errors.New("Validation error")
	}
	return nil
}
